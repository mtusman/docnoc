package pkg

// Collector represents all the issues associated with a particular container (name)
type Collector map[string]*Issues

// MinMaxIssueCollector is used to compare collected statistics vs preconfigured limits
// and consequently create an issue if needed
func (c *Collector) MinMaxIssueCollector(cC ContainerConfig, cSV float64, cSN, cN, cID string) {
	var mapContainerStatNameToType = map[string]MinMaxAllocation{
		"CPUPercentage":    cC.CPU,
		"MemoryPercentage": cC.Memory,
		"BlockWriteMB":     cC.BlockWrite,
		"BlockReadMB":      cC.BlockWrite,
		"NetworkRxMB":      cC.NetworkRx,
		"NetworkTxMB":      cC.NetworkTx,
	}
	issues, ok := (*c)[cN]
	if !ok {
		(*c)[cN] = &Issues{
			containerID:   cID,
			containerName: cN,
		}
		issues = (*c)[cN]
	}
	// if both max and min limits are 0 then don't evaluate anything
	if mapContainerStatNameToType[cSN].Min == 0 && mapContainerStatNameToType[cSN].Max == 0 {
		return
	}
	if cSV < mapContainerStatNameToType[cSN].Min {
		issues.MinMaxUtilisationIssue(cSV, cSN, cID, true)
	} else if cSV > mapContainerStatNameToType[cSN].Max {
		issues.MinMaxUtilisationIssue(cSV, cSN, cID, false)
	}
}
