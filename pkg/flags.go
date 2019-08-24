package pkg

import (
	"os"
	"path"
)

// defaultConfigFileLocation is the deault location where the docnoc config file is stored
var defaultConfigFileLocation string

func init() {
	cwd, _ := os.Getwd()
	defaultConfigFileLocation = path.Join(cwd, "docnoc_config.yaml")
}

// Flags represents all the flags passed by the user
type Flags struct {
	ConfigFile *string
	Detach     *bool
	Timeout    *int
}

// NewFlags returns a new Flag configuration
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
