package pkg

const (
	defaultMinAllocation = 0
	defaultMaxAllocation = 100
)

type (
	MinMaxAllocation struct {
		min int `yaml:"min"`
		max int `yaml:"max"`
	}

	ContainerConfig struct {
		CPU    MinMaxAllocation `yaml:"cpu"`
		Memory MinMaxAllocation `yaml:"memory"`
		Disk   MinMaxAllocation `yaml:"disk"`
		Status struct {
			StatusList []string `yaml:"status"`
		}
	}

	ExcludeContainerConfig struct {
		ContainerConfig
		Exclude []string `yaml:"exclude"`
	}

	Config struct {
		ExcludeContainerConfig `yaml:"default"`
		ContainersConfig       []ContainerConfig `yaml:"containers"`
	}

	DocNocConfig struct {
		Config `yaml:"docnoc"`
	}
)

func NewMinMaxAllocation() MinMaxAllocation {
	return MinMaxAllocation{
		min: defaultMinAllocation,
		max: defaultMaxAllocation,
	}
}

func NewContainerConfig() ContainerConfig {
	return ContainerConfig{
		CPU:    NewMinMaxAllocation(),
		Memory: NewMinMaxAllocation(),
		Disk:   NewMinMaxAllocation(),
	}
}

func NewExcludeContainerConfig() ExcludeContainerConfig {
	return ExcludeContainerConfig{
		ContainerConfig: NewContainerConfig(),
		Exclude:         []string{},
	}
}

func NewConfig() Config {
	return Config{
		ExcludeContainerConfig: NewExcludeContainerConfig(),
		ContainersConfig:       []ContainerConfig{},
	}
}

func NewDocNocConfig() DocNocConfig {
	return DocNocConfig{
		Config: NewConfig(),
	}
}
