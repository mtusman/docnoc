package pkg

type Collector map[string]*Issues

func (c *Collector) CPUIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerName, containerID string) {
	issues, ok := (*c)[containerName]
	if !ok {
		(*c)[containerName] = &Issues{}
		issues = (*c)[containerName]
	}
	if cSS.CPUPercentage < mMA.Min {
		issues.AboveMinUtilisationIssue("CPU", cSS.CPUPercentage, containerID)
	} else if cSS.CPUPercentage > mMA.Max {
		issues.AboveMaxUtilisationIssue("CPU", cSS.CPUPercentage, containerID)
	}
}

func (c *Collector) MemoryIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerName, containerID string) {
	issues, ok := (*c)[containerName]
	if !ok {
		(*c)[containerName] = &Issues{}
		issues = (*c)[containerName]
	}
	if cSS.MemoryPercentage < mMA.Min {
		issues.AboveMinUtilisationIssue("Memory", cSS.MemoryPercentage, containerID)
	} else if cSS.MemoryPercentage > mMA.Max {
		issues.AboveMaxUtilisationIssue("Memory", cSS.MemoryPercentage, containerID)
	}
}

func (c *Collector) BlockReadIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerName, containerID string) {
	issues, ok := (*c)[containerName]
	if !ok {
		(*c)[containerName] = &Issues{}
		issues = (*c)[containerName]
	}
	if cSS.BlockRead < mMA.Min {
		issues.AboveMinUtilisationIssue("BlockRead", cSS.BlockRead, containerID)
	} else if cSS.BlockRead > mMA.Max {
		issues.AboveMaxUtilisationIssue("BlockRead", cSS.BlockRead, containerID)
	}
}

func (c *Collector) BlockWriteIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerName, containerID string) {
	issues, ok := (*c)[containerName]
	if !ok {
		(*c)[containerName] = &Issues{}
		issues = (*c)[containerName]
	}
	if cSS.BlockWrite < mMA.Min {
		issues.AboveMinUtilisationIssue("BlockWrite", cSS.BlockWrite, containerID)
	} else if cSS.BlockWrite > mMA.Max {
		issues.AboveMaxUtilisationIssue("BlockWrite", cSS.BlockWrite, containerID)
	}
}

func (c *Collector) NetworkRxIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerName, containerID string) {
	issues, ok := (*c)[containerName]
	if !ok {
		(*c)[containerName] = &Issues{}
		issues = (*c)[containerName]
	}
	if cSS.NetworkRx < mMA.Min {
		issues.AboveMinUtilisationIssue("NetworkRx", cSS.NetworkRx, containerID)
	} else if cSS.NetworkRx > mMA.Max {
		issues.AboveMaxUtilisationIssue("NetworkRx", cSS.NetworkRx, containerID)
	}
}

func (c *Collector) NetworkTxIssueCollector(mMA *MinMaxAllocation, cSS *ContainerSetStatistics, containerName, containerID string) {
	issues, ok := (*c)[containerName]
	if !ok {
		(*c)[containerName] = &Issues{}
		issues = (*c)[containerName]
	}
	if cSS.NetworkTx < mMA.Min {
		issues.AboveMinUtilisationIssue("NetworkTx", cSS.NetworkTx, containerID)
	} else if cSS.NetworkTx > mMA.Max {
		issues.AboveMaxUtilisationIssue("NetworkTx", cSS.NetworkTx, containerID)
	}
}
