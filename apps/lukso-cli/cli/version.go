package cli

import "github.com/urfave/cli/v2"

func getVersionCommand() *cli.Command {
	VersionCommands := []*cli.Command{
		{
			Name:  "vanguard",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "version", "vanguard"
				return nil
			},
		},
		{
			Name:  "pandora",
			Flags: getLuksoFlags(),
			Action: func(c *cli.Context) error {
				Cmd, Arg = "varsion", "pandora"
				return nil
			},
		},
	}

	version := cli.Command{
		Name:      "version",
		Usage:     "Shows version of all or specific client(s)",
		UsageText: "lukso version [client]\n   [orchestrator, pandora, vanguard, validator, eth2stats-client, lukso-status, all]",
		Flags:     getLuksoFlags(),
		Action: func(c *cli.Context) error {
			Cmd, Arg = "version", "all"
			return nil
		},
		Subcommands: VersionCommands,
	}

	return &version
}
