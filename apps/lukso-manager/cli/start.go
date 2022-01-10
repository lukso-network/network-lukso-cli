package cli

import "github.com/urfave/cli/v2"

func getStartCommand() *cli.Command {
	startCommands := []*cli.Command{
		{
			Name:    "orchestrator",
			Aliases: []string{"lukso-orchestrator"},
			Flags:   getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "start", "lukso-orchestrator"
				return nil
			},
		},
		{
			Name:  "pandora",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "start", "pandora"
				return nil
			},
		},
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "start", "vanguard"
				return nil
			},
		},
		{
			Name:    "validator",
			Aliases: []string{"lukso-validator"},
			Flags:   getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "start", "lukso-validator"
				return nil
			},
		},
		{
			Name:  "all",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "start", "all"
				return nil
			},
		},
	}

	start := cli.Command{
		Name:      "start",
		Usage:     "Starts up all or specific client(s)",
		UsageText: "lukso start [client]\n   [orchestrator, pandora, vanguard, validator, eth2stats-client, lukso-status, all]",
		Flags:     getLuksoFlags(),
		Action: func(c *cli.Context) error {
			Cmd, Arg = "start", "all"
			return nil
		},
		Subcommands: startCommands,
	}

	return &start
}
