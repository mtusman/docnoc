package pkg

import (
	"os"
	"path"
)

var defaultConfigFileLocation string

func init() {
	cwd, _ := os.Getwd()
	defaultConfigFileLocation = path.Join(cwd, "docnoc_config.yaml")
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
