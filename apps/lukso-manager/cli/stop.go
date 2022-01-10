package cli

import "github.com/urfave/cli/v2"

func getStopCommand() *cli.Command {
	stopSubCommands := []*cli.Command{
		{
			Name:    "orchestrator",
			Aliases: []string{"lukso-orchestrator"},
			Flags:   getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "stop", "lukso-orchestrator"
				return nil
			},
		},
		{
			Name:  "pandora",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "stop", "pandora"
				return nil
			},
		},
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "stop", "vanguard"
				return nil
			},
		},
		{
			Name:    "validator",
			Aliases: []string{"lukso-validator"},
			Flags:   getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "stop", "lukso-validator"
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
