package pkg

import (
	"strings"

	"github.com/docker/docker/api/types"
)

var (
	// mapContainerStatNameToIssueName is used to keys of a ContainerStatName struct
	// to a human friendly format
	mapContainerStatNameToIssueName = map[string]string{
		"CPUPercentage":    "CPU",
		"MemoryPercentage": "Memory",
		"BlockWriteMB":     "Block I",
		"BlockReadMB":      "Block O",
		"NetworkRxMB":      "Network I",
		"NetworkTxMB":      "Network O",
	}
	// ContainerStatNamePercs represents issues represented as a percentage
	ContainerStatNamePercs = []string{"CPU", "Memory"}
)

// ContainerSetStatistics represents the statistics collected for a particular container
type ContainerSetStatistics struct {
	CPUPercentage, MemoryPercentage, BlockWriteMB, BlockReadMB, NetworkRxMB, NetworkTxMB float64
}

// calculateCPUPercentUnix is used to calculate CPU usage of a linux container
func calculateCPUPercentUnix(previousCPU, previousSystem uint64, v *types.StatsJSON) float64 {
	var (
		cpuPercent = 0.0
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(v.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
		// calculate the change for the entire system between readings
		systemDelta = float64(v.CPUStats.SystemUsage) - float64(previousSystem)
	)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(v.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}

// calculateCPUPercentWindows is used to calculate CPU usage of a windows container
func calculateCPUPercentWindows(v *types.StatsJSON) float64 {
	// Max number of 100ns intervals between the previous time read and now
	possIntervals := uint64(v.Read.Sub(v.PreRead).Nanoseconds()) // Start with number of ns intervals
	possIntervals /= 100                                         // Convert to number of 100ns intervals
	possIntervals *= uint64(v.NumProcs)                          // Multiple by the number of processors

	// Intervals used
	intervalsUsed := v.CPUStats.CPUUsage.TotalUsage - v.PreCPUStats.CPUUsage.TotalUsage

	// Percentage avoiding divide-by-zero
	if possIntervals > 0 {
		return float64(intervalsUsed) / float64(possIntervals) * 100.0
	}
	return 0.00
}

// calculateBlockIO is used to calculate Block IO of a linux container
func calculateBlockIO(blkio types.BlkioStats) (blkRead uint64, blkWrite uint64) {
	for _, bioEntry := range blkio.IoServiceBytesRecursive {
		switch strings.ToLower(bioEntry.Op) {
		case "read":
			blkRead = blkRead + bioEntry.Value
		case "write":
			blkWrite = blkWrite + bioEntry.Value
		}
	}
	return
}

// calculateNetwork is used to calculate Network RX/TX of both a linux and windows container
func calculateNetwork(network map[string]types.NetworkStats) (float64, float64) {
	var rx, tx float64

	for _, v := range network {
		rx += float64(v.RxBytes)
		tx += float64(v.TxBytes)
	}
	return rx, tx
}
