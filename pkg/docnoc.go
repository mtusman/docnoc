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

func (dN *DocNoc) StartScrubbing() {
	ctx := context.Background()
	containers, err := dN.Client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		fmt.Println("🔥: Can't get a list of containers", err)
	}
	for _, container := range containers {
		cSS := Watcher(ctx, dN.Client, container.ID, false)
		dN.ScrubMinMaxEvaluate(cSS, container.ID)
	}
}

func (dN *DocNoc) ScrubMinMaxEvaluate(cSS *ContainerSetStatistics, containerID string) {
	dN.Collector.CPUIssueCollector(&dN.DocNocConfig.ExcludeContainerConfig.CPU, cSS, containerID)
	dN.Collector.MemoryIssueCollector(&dN.DocNocConfig.ExcludeContainerConfig.Memory, cSS, containerID)
	dN.Collector.BlockWriteIssueCollector(&dN.DocNocConfig.ExcludeContainerConfig.BlockWrite, cSS, containerID)
	dN.Collector.BlockReadIssueCollector(&dN.DocNocConfig.ExcludeContainerConfig.BlockRead, cSS, containerID)
	dN.Collector.NetworkRxIssueCollector(&dN.DocNocConfig.ExcludeContainerConfig.NetworkRx, cSS, containerID)
	dN.Collector.NetworkTxIssueCollector(&dN.DocNocConfig.ExcludeContainerConfig.NetworkTx, cSS, containerID)
}
