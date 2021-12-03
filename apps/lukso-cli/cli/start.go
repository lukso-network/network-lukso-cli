package cli

import "github.com/urfave/cli"

func getStartCommand() cli.Command {
	startCommands := []cli.Command{
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) {
				Cmd, Arg = "start", "vanguard"
			},
		},
		{
			Name:  "pandora",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) {
				Cmd, Arg = "start", "pandora"
			},
		},
	}

	start := cli.Command{
		Name:      "start",
		Usage:     "Starts up all or specific client(s)",
		UsageText: "lukso start [client]\n   [orchestrator, pandora, vanguard, validator, eth2stats-client, lukso-status, all]",
		Flags:     getLuksoFlags(),
		Action: func(c *cli.Context) {
			Cmd, Arg = "start", "all"
		},
		Subcommands: startCommands,
	}

	return start
}
