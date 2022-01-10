package cli

import (
	"github.com/urfave/cli/v2"

	"lukso/apps/lukso-manager/settings"
)

var Placeholder struct {
	settings.Settings
	API        bool
	GUI        bool
	ConfigFile string
	Network    string
	PandoraTag string
}

var luksoFlags []cli.Flag

func getLuksoFlags() []cli.Flag {

	luksoFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "api",
			Usage:       "Starts API server",
			Destination: &Placeholder.API,
		},
		&cli.BoolFlag{
			Name:        "gui",
			Usage:       "Starts GUI",
			Destination: &Placeholder.GUI,
		},
		&cli.StringFlag{
			Name:        "config",
			Value:       "",
			Usage:       "Loads config file",
			TakesFile:   true,
			Destination: &Placeholder.ConfigFile,
		},
		&cli.StringFlag{
			Name:        "network",
			Value:       "",
			Usage:       "Picks which setup to use",
			EnvVars:     []string{"NETWORK"},
			Destination: &Placeholder.Network,
		},
		&cli.StringFlag{
			Name:        "coinbase",
			Usage:       "Sets pandora coinbase. This is public address for block mining rewards",
			Destination: &Placeholder.Coinbase,
		},
		&cli.BoolFlag{
			Name:        "validate",
			Usage:       "Enables validator",
			Destination: &Placeholder.ValidatorEnabled,
		},
		&cli.StringFlag{
			Name:        "orchestrator",
			Aliases:     []string{"orchestrator-tag", "orc-tag"},
			Usage:       "Sets pandora tag version to be used",
			Destination: &Placeholder.PandoraTag,
		},
		&cli.StringFlag{
			Name:        "pandora",
			Aliases:     []string{"pandora-tag", "pan-tag"},
			Usage:       "Sets pandora tag version to be used",
			Destination: &Placeholder.PandoraTag,
		},
	}

	return luksoFlags
}
