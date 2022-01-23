package config

import (
	"os"
	"path/filepath"
)

const (
	// DefaultLogLevel represents the default log level.
	DefaultLogLevel = "info"
)

// DefaultLogFile represents the default K9s log file.
var DefaultLogFile = filepath.Join(os.TempDir(), "lukso.log")

// Flags represents K9s configuration flags.
type Flags struct {
	LogLevel *string
	LogFile  *string
}

// NewFlags returns new configuration flags.
func NewFlags() *Flags {
	return &Flags{
		LogLevel: strPtr(DefaultLogLevel),
		LogFile:  strPtr(DefaultLogFile),
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func strPtr(s string) *string {
	return &s
}
