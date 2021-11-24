package flags

import "github.com/urfave/cli"

func getStartCommand() cli.Command {
	startCommands := []cli.Command{
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) {
				println("Starting Vanguard")
			},
		},
	}

	start := cli.Command{
		Name:      "start",
		Usage:     "Starts up all or specific client(s)",
		UsageText: "lukso start [client]\n   [orchestrator, pandora, vanguard, validator, eth2stats-client, lukso-status, all]",
		Flags:     getLuksoFlags(),
		Action: func(c *cli.Context) {
			println("Starting all")
		},
		Subcommands: startCommands,
	}

	return start
}
