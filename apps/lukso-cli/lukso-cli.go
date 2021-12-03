package main

import (
	"lukso-cli/config"
	"lukso-cli/cli"
	"lukso-cli/runner"
)

var LuksoSettings config.LuksoValues

func main() {
	cli.InitFlags()

	// Build Settings

	//Load from default first
	config.LoadDefaults(&LuksoSettings)

	//Overwrite with config
	if cli.FlagValues.Config != "" {
		config.LoadConfig(&LuksoSettings, cli.FlagValues.Config)
	}

	//Overwrite with flags
	cli.LoadFlags(&LuksoSettings)

	// RUN

	runner.Action(cli.Cmd, cli.Arg, &LuksoSettings)

}
