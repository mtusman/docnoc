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
	}

	DocNocConfig struct {
		Config `yaml:"docnoc"`
	}
)

func NewMinMaxAllocation() MinMaxAllocation {
	return MinMaxAllocation{
		Min: defaultMinAllocation,
		Max: defaultMaxAllocation,
	}
}

func NewDefaultContainerConfig() ContainerConfig {
	return ContainerConfig{
		CPU:        NewMinMaxAllocation(),
		Memory:     NewMinMaxAllocation(),
		Disk:       NewMinMaxAllocation(),
		BlockRead:  NewMinMaxAllocation(),
		BlockWrite: NewMinMaxAllocation(),
		NetworkRx:  NewMinMaxAllocation(),
		NetworkTx:  NewMinMaxAllocation(),
	}
}

func NewConfig() Config {
	return Config{
		DefaultContainerConfig: NewDefaultContainerConfig(),
		ContainersConfig:       map[string]ContainerConfig{},
		Exclude:                []string{},
	}
}

func NewDocNocConfig() DocNocConfig {
	return DocNocConfig{
		Config: NewConfig(),
	}
}
