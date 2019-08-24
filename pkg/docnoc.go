package pkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-yaml/yaml"
)

type DocNoc struct {
	Client       *client.Client
	DocNocConfig *DocNocConfig
	Flags        *Flags
	Collector    Collector
	Context      context.Context
}

func NewDocNoc(flags *Flags) *DocNoc {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.39"))
	if err != nil {
		fmt.Println("🔥: Failed to create docker client")
		os.Exit(1)
	}

	var f []byte
	if flags.ConfigFile != nil {
		cwd, _ := os.Getwd()
		f, err = ioutil.ReadFile(path.Join(cwd, *flags.ConfigFile))
	} else {
		f, err = ioutil.ReadFile(defaultConfigFileLocation)
	}

	if err != nil {
		fmt.Println("🔥: Unable to read docnoc config file")
		os.Exit(1)
	}

	cfg := NewDocNocConfig()
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		fmt.Println("🔥: Can't unmarshall docnoc config file", err)
		os.Exit(1)
	}
	return &DocNoc{
		Client:       cli,
		DocNocConfig: &cfg,
		Flags:        flags,
		Collector:    Collector{},
		Context:      context.Background(),
	}
}

func (dN *DocNoc) StartScrubbingDefault() {
	containers, err := dN.Client.ContainerList(dN.Context, types.ContainerListOptions{})
	if err != nil {
		fmt.Println("🔥: Failed to get list of containers", err)
		os.Exit(1)
	}

	if dN.DocNocConfig.Config.SlackWebhook != "" {
		PostInitSlackMessage(dN.DocNocConfig.Config.SlackWebhook)
	}

	for _, container := range containers {
		cN := container.Names[0][1:]
		inExclude := containerNameInExclude(cN, dN.DocNocConfig.Exclude)
		_, inSeperateConfig := dN.DocNocConfig.ContainersConfig[cN]
		if inSeperateConfig {
			cSS := Watcher(dN.Context, dN.Client, container.ID, false)
			scrubMinMaxEvaluate(dN.Collector, dN.DocNocConfig.ContainersConfig[cN], cSS, cN, container.ID)
		} else if !inExclude {
			cSS := Watcher(dN.Context, dN.Client, container.ID, false)
			scrubMinMaxEvaluate(dN.Collector, dN.DocNocConfig.DefaultContainerConfig, cSS, cN, container.ID)
		}
	}
}

func (dN *DocNoc) ProcessReport() {
	PrintTitle("default")
	for key, issues := range dN.Collector {
		inExclude := containerNameInExclude(key, dN.DocNocConfig.Exclude)
		if !inExclude {
			dN.processReportForApp(key, issues, dN.DocNocConfig.DefaultContainerConfig)
		}
	}

	for key, cC := range dN.DocNocConfig.ContainersConfig {
		issues, ok := dN.Collector[key]
		if ok {
			PrintTitle(key)
			dN.processReportForApp(key, issues, cC)
		}
	}
}

func (dN *DocNoc) processReportForApp(key string, issues *Issues, cC ContainerConfig) {
	numErrs := len((*issues).IssuesList)
	PrintContainerName(key, numErrs)
	if numErrs != 0 {
		PrintIssuesList(dN, key, issues.containerID, issues.IssuesList)
		if issues.ActionTaken == false {
			if cC.Action == "stop" {
				err := dN.Client.ContainerStop(dN.Context, issues.containerID, nil)
				if err != nil {
					PostActionMessage(dN.DocNocConfig.Config.SlackWebhook, key, issues.containerID, "stop", true)
					fmt.Println(err)
				} else {
					PostActionMessage(dN.DocNocConfig.Config.SlackWebhook, key, issues.containerID, "Stopped", false)
				}
			} else if cC.Action == "restart" {
				err := dN.Client.ContainerRestart(dN.Context, issues.containerID, nil)
				if err != nil {
					PostActionMessage(dN.DocNocConfig.Config.SlackWebhook, key, issues.containerID, "restart", true)
					fmt.Println(err)
				} else {
					PostActionMessage(dN.DocNocConfig.Config.SlackWebhook, key, issues.containerID, "Restarted", false)
				}
			}
		}
	}

}

func containerNameInExclude(name string, Exclude []string) bool {
	for _, v := range Exclude {
		if v == name {
			return true
		}
	}
	return false
}

func scrubMinMaxEvaluate(clctr Collector, cC ContainerConfig, cSS *ContainerSetStatistics, containerName, containerID string) {
	v := reflect.ValueOf(*cSS)
	typeOfcSS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		clctr.MinMaxIssueCollector(cC, v.Field(i).Interface().(float64), typeOfcSS.Field(i).Name, containerName, containerID)
	}

}
