package pkg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ExcludeWatcher(ecc *ExcludeContainerConfig, cli *client.Client) {
	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		fmt.Println("🔥: Can't get a list of containers", err)
	}
	for _, container := range containers {
		stat, err := cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			fmt.Println(fmt.Sprintf("Stats collection failed for container with ID %s", container.ID))
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
			fmt.Println(memUsage, memLimit)
			memPerc = (memUsage / memLimit) * 100
		} else {
			cpuPerc = calculateCPUPercentWindows(&containerStatsJSON)
			blkRead = containerStatsJSON.StorageStats.ReadSizeBytes
			blkWrite = containerStatsJSON.StorageStats.WriteSizeBytes
		}
		netRx, netTx := calculateNetwork(containerStatsJSON.Networks)
		css := ContainerSetStatistics{
			CPUPercentage:    cpuPerc,
			MemoryPercentage: memPerc,
			BlockRead:        float64(blkRead),
			BlockWrite:       float64(blkWrite),
			NetworkRx:        netRx,
			NetworkTx:        netTx,
		}
		fmt.Println(css)

	}
}
