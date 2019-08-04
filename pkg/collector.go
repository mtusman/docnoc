package pkg

type Collector map[string]*Issues

func (c *Collector) CPUIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerID string) {
	_, ok := (*c)[containerID]
	if !ok {
		(*c)[containerID] = &Issues{}
	}
	if cSS.CPUPercentage < mMA.Min {
		(*c)[containerID].AboveMinUtilisationIssue("CPU", cSS.CPUPercentage)
	} else if cSS.CPUPercentage > mMA.Max {
		(*c)[containerID].AboveMaxUtilisationIssue("CPU", cSS.CPUPercentage)
	}
}

func (c *Collector) MemoryIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerID string) {
	_, ok := (*c)[containerID]
	if !ok {
		(*c)[containerID] = &Issues{}
	}
	if cSS.MemoryPercentage < mMA.Min {
		(*c)[containerID].AboveMinUtilisationIssue("Memory", cSS.MemoryPercentage)
	} else if cSS.MemoryPercentage > mMA.Max {
		(*c)[containerID].AboveMaxUtilisationIssue("Memory", cSS.MemoryPercentage)
	}
}

func (c *Collector) BlockReadIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerID string) {
	_, ok := (*c)[containerID]
	if !ok {
		(*c)[containerID] = &Issues{}
	}
	if cSS.BlockRead < mMA.Min {
		(*c)[containerID].AboveMinUtilisationIssue("BlockRead", cSS.BlockRead)
	} else if cSS.BlockRead > mMA.Max {
		(*c)[containerID].AboveMaxUtilisationIssue("BlockRead", cSS.BlockRead)
	}
}

func (c *Collector) BlockWriteIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerID string) {
	_, ok := (*c)[containerID]
	if !ok {
		(*c)[containerID] = &Issues{}
	}
	if cSS.BlockWrite < mMA.Min {
		(*c)[containerID].AboveMinUtilisationIssue("BlockWrite", cSS.BlockWrite)
	} else if cSS.BlockWrite > mMA.Max {
		(*c)[containerID].AboveMaxUtilisationIssue("BlockWrite", cSS.BlockWrite)
	}
}

func (c *Collector) NetworkRxIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerID string) {
	_, ok := (*c)[containerID]
	if !ok {
		(*c)[containerID] = &Issues{}
	}
	if cSS.NetworkRx < mMA.Min {
		(*c)[containerID].AboveMinUtilisationIssue("NetworkRx", cSS.NetworkRx)
	} else if cSS.NetworkRx > mMA.Max {
		(*c)[containerID].AboveMaxUtilisationIssue("NetworkRx", cSS.NetworkRx)
	}
}

func (c *Collector) NetworkTxIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerID string) {
	_, ok := (*c)[containerID]
	if !ok {
		(*c)[containerID] = &Issues{}
	}
	if cSS.NetworkTx < mMA.Min {
		(*c)[containerID].AboveMinUtilisationIssue("NetworkTx", cSS.NetworkTx)
	} else if cSS.NetworkTx > mMA.Max {
		(*c)[containerID].AboveMaxUtilisationIssue("NetworkTx", cSS.NetworkTx)
	}
}
