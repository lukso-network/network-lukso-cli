package main

import (
	"lukso-cli/config"
	"lukso-cli/flags"
	"lukso-cli/runner"
)

var LuksoSettings config.LuksoValues

func main() {
	flags.InitFlags()

	// Build Settings

	//Load from default first
	config.LoadDefaults(&LuksoSettings)

	//Overwrite with config
	if flags.FlagValues.Config != "" {
		config.LoadConfig(&LuksoSettings, flags.FlagValues.Config)
	}

	//Overwrite with flags
	flags.LoadFlags(&LuksoSettings)

	// RUN

	runner.Action(flags.Cmd, flags.Arg, &LuksoSettings)

}
