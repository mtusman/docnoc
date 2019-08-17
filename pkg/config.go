package pkg

const (
	defaultMinAllocation = 0
	defaultMaxAllocation = 100
)

type (
	MinMaxAllocation struct {
		Min float64 `yaml:"min"`
		Max float64 `yaml:"max"`
	}

	ContainerConfig struct {
		CPU        MinMaxAllocation `yaml:"cpu"`
		Memory     MinMaxAllocation `yaml:"memory"`
		Disk       MinMaxAllocation `yaml:"disk"`
		BlockWrite MinMaxAllocation `yaml:"block_write"`
		BlockRead  MinMaxAllocation `yaml:"block_read"`
		NetworkRx  MinMaxAllocation `yaml:"network_rx"`
		NetworkTx  MinMaxAllocation `yaml:"network_tx"`
		Status     []string         `yaml:"status"`
	}

	Config struct {
		DefaultContainerConfig ContainerConfig            `yaml:"default"`
		ContainersConfig       map[string]ContainerConfig `yaml:"containers"`
		Exclude                []string                   `yaml:"exclude"`
		SlackWebhook           string                     `yaml:"slack_webhook"`
	}

	DocNocConfig struct {
		Config `yaml:"docnoc"`
	}
)

func newMinMaxAllocation() MinMaxAllocation {
	return MinMaxAllocation{
		Min: defaultMinAllocation,
		Max: defaultMaxAllocation,
	}
}

func newDefaultContainerConfig() ContainerConfig {
	return ContainerConfig{
		CPU:        newMinMaxAllocation(),
		Memory:     newMinMaxAllocation(),
		Disk:       newMinMaxAllocation(),
		BlockRead:  newMinMaxAllocation(),
		BlockWrite: newMinMaxAllocation(),
		NetworkRx:  newMinMaxAllocation(),
		NetworkTx:  newMinMaxAllocation(),
	}
}

func newConfig() Config {
	return Config{
		DefaultContainerConfig: newDefaultContainerConfig(),
		ContainersConfig:       map[string]ContainerConfig{},
		Exclude:                []string{},
	}
}

func NewDocNocConfig() DocNocConfig {
	return DocNocConfig{
		Config: newConfig(),
	}
}
