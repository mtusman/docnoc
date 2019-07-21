package pkg

const (
	defaultMinAllocation = 0
	defaultMaxAllocation = 100
)

type (
	MinMaxAllocation struct {
		Min int `yaml:"min"`
		Max int `yaml:"max"`
	}

	ContainerConfig struct {
		CPU    MinMaxAllocation `yaml:"cpu"`
		Memory MinMaxAllocation `yaml:"memory"`
		Disk   MinMaxAllocation `yaml:"disk"`
		Status []string         `yaml:"status"`
	}

	ExcludeContainerConfig struct {
		CPU     MinMaxAllocation `yaml:"cpu"`
		Memory  MinMaxAllocation `yaml:"memory"`
		Disk    MinMaxAllocation `yaml:"disk"`
		Status  []string         `yaml:"status"`
		Exclude []string         `yaml:"exclude"`
	}

	Config struct {
		ExcludeContainerConfig `yaml:"default"`
		ContainersConfig       map[string]ContainerConfig `yaml:"containers"`
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

func NewExcludeContainerConfig() ExcludeContainerConfig {
	return ExcludeContainerConfig{
		CPU:    NewMinMaxAllocation(),
		Memory: NewMinMaxAllocation(),
		Disk:   NewMinMaxAllocation(),
	}
}

func NewConfig() Config {
	return Config{
		ExcludeContainerConfig: NewExcludeContainerConfig(),
		ContainersConfig:       map[string]ContainerConfig{},
	}
}

func NewDocNocConfig() DocNocConfig {
	return DocNocConfig{
		Config: NewConfig(),
	}
}
