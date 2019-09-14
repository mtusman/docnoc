package pkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-yaml/yaml"
)

// DocNoc represents a docker sanitizer
type DocNoc struct {
	Client       *client.Client
	DocNocConfig *DocNocConfig
	Flags        *Flags
	Collector    Collector
	Context      context.Context
}

// NewDocNoc returns a new DocNoc configuration
func NewDocNoc(flags *Flags) *DocNoc {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.39"))
	if err != nil {
		fmt.Println("🔥: Failed to create docker client")
		os.Exit(1)
	}

	var f []byte
	if flags.ConfigFile != nil {
		f, err = ioutil.ReadFile(*flags.ConfigFile)
	} else {
		f, err = ioutil.ReadFile(defaultConfigFileLocation)
	}

	if err != nil {
		fmt.Println("🔥: Unable to read docnoc config file", err)
		os.Exit(1)
	}

	cfg := NewDocNocConfig()
	ctx := context.Background()

	// if docnoc is running inside a container, then add its name into exclude
	out, err := exec.Command("cat /etc/hostname").Output()
	if err == nil {
		// docnoc is running inside container
		containerInspect, err := cli.ContainerInspect(ctx, string(out))
		if err != nil {
			fmt.Println("🔥:Failed to get name of host container", err)
			os.Exit(1)
		}
		cfg.Exclude = append(cfg.Exclude, containerInspect.Name)
	}
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
		Context:      ctx,
	}
}

// StartScrubbingDefault goes through each container specified in the config file,
// collects its statistics, and collects issues if those statistics fail to meet
// predefined conditions
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

// ProcessReport is used to output all the issues to both the terminal as well as send
// issues and actions reports to a slack webhook (if specified)
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

// processReportForApp process a report for a single issue
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

// containerNameInExclude returns if we don't need to sanitize a certain container
func containerNameInExclude(name string, Exclude []string) bool {
	for _, v := range Exclude {
		if v == name {
			return true
		}
	}
	return false
}

// scrubMinMaxEvaluate goes throught each attribute and compares it with preconfigured limits set in
// the docnoc file
func scrubMinMaxEvaluate(clctr Collector, cC ContainerConfig, cSS *ContainerSetStatistics, containerName, containerID string) {
	v := reflect.ValueOf(*cSS)
	typeOfcSS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		clctr.MinMaxIssueCollector(cC, v.Field(i).Interface().(float64), typeOfcSS.Field(i).Name, containerName, containerID)
	}

}
