package cli

import "github.com/urfave/cli/v2"

func getStopCommand() *cli.Command {
	stopSubCommands := []*cli.Command{
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "stop", "vanguard"
				return nil
			},
		},
	}

	stop := cli.Command{
		Name:      "stop",
		Usage:     "Stop up all or specific client(s)",
		UsageText: "lukso stop [client]\n   [orchestrator, pandora, vanguard, validator, eth2stats-client, lukso-status, all]",
		Flags:     getLuksoFlags(),
		Action: func(c *cli.Context) error {
			Cmd, Arg = "stop", "all"
			return nil
		},
		Subcommands: stopSubCommands,
	}

	return &stop
}
