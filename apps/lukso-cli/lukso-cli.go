package main

import (
	"lukso-cli/config"
	"lukso-cli/flags"
)

var LuksoSettings config.LuksoValues

func main() {
	flags.InitFlags()

	// Build Settings

	//Load from default first
	config.LoadDefaults(&LuksoSettings)

	//Overwrite with config
	if flags.FlagValues.Config != "" {
		println("Config loaded")
		config.LoadConfig(&LuksoSettings, flags.FlagValues.Config)
	}

	//Overwrite with flags
	flags.LoadFlags(&LuksoSettings)

	//check
	println(LuksoSettings.Network)
	println(LuksoSettings.Orchestrator.Verbosity)
	println(LuksoSettings.Pandora.Verbosity)
}
