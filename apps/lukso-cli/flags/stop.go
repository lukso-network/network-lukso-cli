package flags

import "github.com/urfave/cli"

func getStopCommand() cli.Command {
	stopSubCommands := []cli.Command{
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) {
				println("Stopping Vanguard")
			},
		},
	}

	stop := cli.Command{
		Name:      "stop",
		Usage:     "Stop up all or specific client(s)",
		UsageText: "lukso stop [client]\n   [orchestrator, pandora, vanguard, validator, eth2stats-client, lukso-status, all]",
		Flags:     getLuksoFlags(),
		Action: func(c *cli.Context) {
			println("Stopping all")
		},
		Subcommands: stopSubCommands,
	}

	return stop
}
