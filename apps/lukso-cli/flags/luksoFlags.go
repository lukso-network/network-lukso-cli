package flags

import (
	"github.com/urfave/cli"

	"lukso-cli/config"
)

var FlagValues struct {
	config.LuksoValues
	Config      string
	l15_prod    bool
	l15_staging bool
	l15_dev     bool
}

var luksoFlags []cli.Flag

func getLuksoFlags() []cli.Flag {
	luksoFlags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Value:       "",
			Usage:       "Loads config file",
			Destination: &FlagValues.Config,
		},
		cli.StringFlag{
			Name:        "network",
			Value:       "",
			Usage:       "Picks which setup to use",
			Destination: &FlagValues.Network,
		},
		cli.BoolFlag{
			Name:        "l15-prod",
			Usage:       "Alias for --network l15-prod",
			Destination: &FlagValues.l15_prod,
		},
		cli.BoolFlag{
			Name:        "l15-staging",
			Usage:       "Alias for --network l15-staging",
			Destination: &FlagValues.l15_staging,
		},
		cli.BoolFlag{
			Name:        "l15-dev",
			Usage:       "Alias for --network l15-dev",
			Destination: &FlagValues.l15_dev,
		},
	}
	return luksoFlags
}
