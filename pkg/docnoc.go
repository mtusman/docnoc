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
	Collector    *Collector
}

func NewDocNoc(flags *Flags) *DocNoc {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.39"))
	if err != nil {
		fmt.Println("🔥: Can't connect to docker client")
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
		fmt.Println("🔥: Unable to read config file")
		os.Exit(1)
	}

	cfg := NewDocNocConfig()
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		fmt.Println("🔥: Can't unmarshall yaml file", err)
	}
	return &DocNoc{
		Client:       cli,
		DocNocConfig: &cfg,
		Flags:        flags,
		Collector:    &Collector{},
	}
}

func (dN *DocNoc) StartScrubbingDefault() {
	ctx := context.Background()
	containers, err := dN.Client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		fmt.Println("🔥: Can't get a list of containers", err)
	}
	for _, container := range containers {
		inExclude := containerNameInExclude(container.Names[0], dN.DocNocConfig.Exclude)
		if !inExclude {
			cSS := Watcher(ctx, dN.Client, container.ID, false)
			scrubMinMaxEvaluate(dN.Collector, dN.DocNocConfig.DefaultContainerConfig, cSS, container.Names[0], container.ID)
		}
	}
	dN.OutputResultsForSection("default")
}

func containerNameInExclude(name string, Exclude []string) bool {
	for _, v := range Exclude {
		if v == name {
			return true
		}
	}
	return false
}

func scrubMinMaxEvaluate(clctr *Collector, cC ContainerConfig, cSS *ContainerSetStatistics, containerName, containerID string) {
	v := reflect.ValueOf(*cSS)
	typeOfcSS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		clctr.MinMaxIssueCollector(cC, v.Field(i).Interface().(float64), typeOfcSS.Field(i).Name, containerName, containerID)
	}

}

func (dN *DocNoc) OutputResultsForSection(section string) {
	printTitle(section)
	for key, issues := range *(dN.Collector) {
		issLen := len(*issues)
		printContainerName(key, issLen)

		if issLen != 0 {
			for containerID, issueList := range *issues {
				printContainerID(containerID)
				printIssuesList(issueList)
			}
		}
	}
}
