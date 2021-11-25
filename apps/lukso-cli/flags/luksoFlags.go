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

	var DefaultValues config.LuksoValues
	config.LoadDefaults(&DefaultValues)

	luksoFlags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Value:       "",
			Usage:       "Loads config file",
			TakesFile:   true,
			Destination: &FlagValues.Config,
		},
		cli.StringFlag{
			Name:        "network",
			Value:       "",
			Usage:       "Picks which setup to use default: " + DefaultValues.Network,
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
		cli.BoolFlag{
			Name:        "validate",
			Usage:       "Starts validator",
			Destination: &FlagValues.Validate,
		},
		cli.StringFlag{
			Name:        "coinbase",
			Usage:       "Sets pandora coinbase. This is public address for block mining rewards (default = first account created) (default: \"0\")",
			Destination: &FlagValues.Coinbase,
		},
		cli.StringFlag{
			Name:        "node-name",
			Usage:       "Name of node that's shown on pandora stats and vanguard stats",
			Destination: &FlagValues.NodeName,
		},
		cli.StringFlag{
			Name:        "logsdir",
			Usage:       "Sets the logs path",
			Destination: &FlagValues.LogsDir,
		},
		cli.StringFlag{
			Name:        "datadir",
			Usage:       "Sets datadir path",
			Destination: &FlagValues.DataDir,
		},
		cli.StringFlag{
			Name:        "orchestrator-verbosity, orc-verbosity",
			Usage:       "Sets orchestrator logging depth",
			Destination: &FlagValues.Orchestrator.Verbosity,
		},
	}
	return luksoFlags
}
