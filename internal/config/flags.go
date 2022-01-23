package config

import (
	"os"
	"path/filepath"
)

const (
	// DefaultLogLevel represents the default log level.
	DefaultLogLevel   = "info"
	DefaultConfigFile = ".lukso.yaml"
)

// DefaultLogFile represents the default K9s log file.
var DefaultLogFile = filepath.Join(os.TempDir(), "lukso.log")

// Flags represents K9s configuration flags.
type Flags struct {
	LogLevel   string
	LogFile    string
	ConfigFile string
}

// NewFlags returns new configuration flags.
func NewFlags() *Flags {
	return &Flags{
		LogLevel:   DefaultLogLevel,
		LogFile:    DefaultLogFile,
		ConfigFile: DefaultConfigFile,
	}
}
