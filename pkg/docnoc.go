package pkg

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"

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
			dN.ScrubMinMaxEvaluate(cSS, container.Names[0], container.ID)
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

func (dN *DocNoc) ScrubMinMaxEvaluate(cSS *ContainerSetStatistics, containerName, containerID string) {
	dN.Collector.CPUIssueCollector(&dN.DocNocConfig.DefaultContainerConfig.CPU, cSS, containerName, containerID)
	dN.Collector.MemoryIssueCollector(&dN.DocNocConfig.DefaultContainerConfig.Memory, cSS, containerName, containerID)
	dN.Collector.BlockWriteIssueCollector(&dN.DocNocConfig.DefaultContainerConfig.BlockWrite, cSS, containerName, containerID)
	dN.Collector.BlockReadIssueCollector(&dN.DocNocConfig.DefaultContainerConfig.BlockRead, cSS, containerName, containerID)
	dN.Collector.NetworkRxIssueCollector(&dN.DocNocConfig.DefaultContainerConfig.NetworkRx, cSS, containerName, containerID)
	dN.Collector.NetworkTxIssueCollector(&dN.DocNocConfig.DefaultContainerConfig.NetworkTx, cSS, containerName, containerID)
}

func (dN *DocNoc) OutputResultsForSection(section string) {
	printTitle(section)
	for key, issues := range *(dN.Collector) {
		issLen := len(*issues)
		printContainerName(key, issLen)

		if issLen != 0 {
			for containerID, issueList := range *issues {
				printContainerID(containerID)
				printIssues(issueList)
			}
		}
	}
}
