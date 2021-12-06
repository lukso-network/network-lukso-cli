package main

import (
	"lukso-cli/cli"
	"lukso-cli/config"
	"lukso-cli/runner"
)

var LuksoSettings config.LuksoValues

func main() {

	// Initialize CLI, read values from flags
	cli.Init()

	// Build Settings

	//Load from default first
	config.LoadDefaults(&LuksoSettings)

	//Overwrite with config
	if cli.FlagValues.Config != "" {
		config.LoadConfig(&LuksoSettings, cli.FlagValues.Config)
	}

	//Overwrite with flags
	cli.LoadFlags(&LuksoSettings)

	//Download binaries if missing

	if LuksoSettings.Orchestrator.Tag != "" {
		//runner.Action("download", "orchestrator", &LuksoSettings)
	}

	// RUN
	runner.Action(cli.Cmd, cli.Arg, &LuksoSettings)

	if cli.FlagValues.ApiServer || cli.FlagValues.GUI {
		//Start API server
		//HTTPApi.Start()
	}

	if cli.FlagValues.GUI {
		// Start Frontend Server
		// gui.Start()
	}

}
