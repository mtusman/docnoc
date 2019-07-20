package pkg

import "os"

var defaultConfigFileLocation string

func init() {
	cwd, _ := os.Getwd()
	defaultConfigFileLocation = cwd + "docnoc_config.yaml"
}

type Flags struct {
	ConfigFile *string
	Detach     *bool
	Timeout    *int
}

func NewFlags() *Flags {
	return &Flags{
		ConfigFile: stringPointer(defaultConfigFileLocation),
		Detach:     boolPointer(false),
	}
}

func stringPointer(s string) *string {
	return &s
}

func boolPointer(b bool) *bool {
	return &b
}