package pkg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Watcher(ctx context.Context, cli *client.Client, containerID string, stream bool) *ContainerSetStatistics {
	stat, err := cli.ContainerStats(ctx, containerID, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("Stats collection failed for container with ID %s", containerID))
	}
	var containerStats types.Stats
	err = json.NewDecoder(stat.Body).Decode(&containerStats)
	containerStatsJSON := types.StatsJSON{Stats: containerStats}
	daemonOSType := stat.OSType

	var (
		cpuPerc  = 0.0
		memPerc  = 0.0
		blkRead  uint64
		blkWrite uint64
	)

	if daemonOSType != "windows" {
		previousCPU := containerStatsJSON.PreCPUStats.CPUUsage.TotalUsage
		previousSystem := containerStatsJSON.PreCPUStats.SystemUsage
		cpuPerc = calculateCPUPercentUnix(previousCPU, previousSystem, &containerStatsJSON)
		blkRead, blkWrite = calculateBlockIO(containerStatsJSON.BlkioStats)
		memUsage := float64(containerStatsJSON.MemoryStats.Usage)
		memLimit := float64(containerStatsJSON.MemoryStats.Limit)
		memPerc = (memUsage / memLimit) * 100
	} else {
		cpuPerc = calculateCPUPercentWindows(&containerStatsJSON)
		blkRead = containerStatsJSON.StorageStats.ReadSizeBytes
		blkWrite = containerStatsJSON.StorageStats.WriteSizeBytes
	}
	netRx, netTx := calculateNetwork(containerStatsJSON.Networks)
	return &ContainerSetStatistics{
		CPUPercentage:    cpuPerc,
		MemoryPercentage: memPerc,
		BlockReadMB:      float64(blkRead / 100000),
		BlockWriteMB:     float64(blkWrite / 100000),
		NetworkRxMB:      netRx,
		NetworkTxMB:      netTx,
	}
}
