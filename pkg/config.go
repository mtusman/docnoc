package pkg

import (
	"math"
)

const defaultMinAllocation = 0

var defaultMaxAllocation = math.Inf(0)

type (
	// MinMaxAllocation represents a min max preconfigured constriction
	MinMaxAllocation struct {
		Min float64 `yaml:"min"`
		Max float64 `yaml:"max"`
	}

	// ContainerConfig represents preconfigured constriction for a container
	ContainerConfig struct {
		CPU        MinMaxAllocation `yaml:"cpu"`
		Memory     MinMaxAllocation `yaml:"memory"`
		BlockWrite MinMaxAllocation `yaml:"block_write"`
		BlockRead  MinMaxAllocation `yaml:"block_read"`
		NetworkRx  MinMaxAllocation `yaml:"network_rx"`
		NetworkTx  MinMaxAllocation `yaml:"network_tx"`
		Action     string           `yaml:"action"`
	}

	// Config represents all the  preconfigured constrictions specified
	// in the docnoc file
	Config struct {
		DefaultContainerConfig ContainerConfig            `yaml:"default"`
		ContainersConfig       map[string]ContainerConfig `yaml:"containers"`
		Exclude                []string                   `yaml:"exclude"`
		SlackWebhook           string                     `yaml:"slack_webhook"`
	}

	// DocNocConfig a structural representation of the docnoc file
	DocNocConfig struct {
		Config `yaml:"docnoc"`
	}
)

// newMinMaxAllocation represents a MinMaxAllocation configuration
func newMinMaxAllocation() MinMaxAllocation {
	return MinMaxAllocation{
		Min: defaultMinAllocation,
		Max: defaultMaxAllocation,
	}
}

// newDefaultContainerConfig represents a ContainerConfig configuration
// for containers not specified or excluded in the docnoc file
func newDefaultContainerConfig() ContainerConfig {
	return ContainerConfig{
		CPU:        newMinMaxAllocation(),
		Memory:     newMinMaxAllocation(),
		BlockRead:  newMinMaxAllocation(),
		BlockWrite: newMinMaxAllocation(),
		NetworkRx:  newMinMaxAllocation(),
		NetworkTx:  newMinMaxAllocation(),
	}
}

// newConfig represents a new Config configuration
func newConfig() Config {
	return Config{
		DefaultContainerConfig: newDefaultContainerConfig(),
		ContainersConfig:       map[string]ContainerConfig{},
		Exclude:                []string{},
	}
}

// NewDocNocConfig represents a new DocNocCon configuration
func NewDocNocConfig() DocNocConfig {
	return DocNocConfig{
		Config: newConfig(),
	}
}
