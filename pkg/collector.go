package pkg

type Collector map[string]*Issues

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
			containerID: cID,
		}
		issues = (*c)[cN]
	}
	if mapContainerStatNameToType[cSN].Min == 0 && mapContainerStatNameToType[cSN].Max == 0 {
		return
	}
	if cSV < mapContainerStatNameToType[cSN].Min {
		issues.MinMaxUtilisationIssue(cSV, cSN, cID, true)
	} else if cSV > mapContainerStatNameToType[cSN].Max {
		issues.MinMaxUtilisationIssue(cSV, cSN, cID, false)
	}
}
