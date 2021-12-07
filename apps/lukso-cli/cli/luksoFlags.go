package cli

import (
	"strconv"

	"github.com/urfave/cli/v2"

	"lukso-cli/config"
)

var FlagValues struct {
	config.LuksoValues
	Config      string
	l15_prod    bool
	l15_staging bool
	l15_dev     bool
	GUI         bool
	ApiServer   bool
}

var luksoFlags []cli.Flag

func getLuksoFlags() []cli.Flag {

	var DefaultValues config.LuksoValues
	config.LoadDefaults(&DefaultValues)

	luksoFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "api",
			Usage:       "Starts API server",
			Destination: &FlagValues.ApiServer,
		},
		&cli.BoolFlag{
			Name:        "gui",
			Usage:       "Starts GUI",
			Destination: &FlagValues.GUI,
		},
		&cli.StringFlag{
			Name:        "config",
			Value:       "",
			Usage:       "Loads config file",
			TakesFile:   true,
			Destination: &FlagValues.Config,
		},
		&cli.BoolFlag{
			Name:        "l15-prod",
			Usage:       "Alias for --network l15-prod",
			Destination: &FlagValues.l15_prod,
		},
		&cli.BoolFlag{
			Name:        "l15-staging",
			Usage:       "Alias for --network l15-staging",
			Destination: &FlagValues.l15_staging,
		},
		&cli.BoolFlag{
			Name:        "l15-dev",
			Usage:       "Alias for --network l15-dev",
			Destination: &FlagValues.l15_dev,
		},
		&cli.BoolFlag{
			Name:        "force",
			Usage:       "Enables force mode",
			Destination: &FlagValues.Force,
		},
		&cli.StringFlag{
			Name:        "network",
			Value:       "",
			Usage:       "Picks which setup to use default: " + DefaultValues.Network,
			Destination: &FlagValues.Network,
			EnvVars:     []string{"NETWORK"},
		},
		&cli.BoolFlag{
			Name:        "validate",
			Usage:       "Starts validator",
			Destination: &FlagValues.Validate,
		},
		&cli.StringFlag{
			Name:        "coinbase",
			Usage:       "Sets pandora coinbase. This is public address for block mining rewards",
			Destination: &FlagValues.Coinbase,
		},
		&cli.StringFlag{
			Name:        "node-name",
			Usage:       "Name of node that's shown on pandora stats and vanguard stats",
			Destination: &FlagValues.NodeName,
		},
		&cli.StringFlag{
			Name:        "datadir",
			Usage:       "Sets datadir path",
			Destination: &FlagValues.DataDir,
		},
		&cli.StringFlag{
			Name:        "logsdir",
			Usage:       "Sets the logs path",
			Destination: &FlagValues.LogsDir,
		},
		&cli.StringFlag{
			Name:        "keys-dir",
			Usage:       "Sets directory of lukso-deposit-cli keys (can be used with \"keygen\" or \"wallet\")",
			Destination: &FlagValues.LogsDir,
		},
		&cli.StringFlag{
			Name:        "keys-password-file",
			Usage:       "Sets path to lukso-deposit-cli keys (can be used with \"keygen\" or \"wallet\")",
			Destination: &FlagValues.LogsDir,
		},
		&cli.StringFlag{
			Name:        "wallet-dir",
			Usage:       "Sets directory of lukso-validator wallet",
			Destination: &FlagValues.LogsDir,
		},
		&cli.StringFlag{
			Name:        "wallet-password-file",
			Usage:       "Password for lukso-validator",
			Destination: &FlagValues.LogsDir,
		},
		&cli.StringFlag{
			Name:        "orchestrator-tag",
			Aliases:     []string{"orc-tag", "orchestrator"},
			Usage:       "download and set orchestrator to given tag",
			Destination: &FlagValues.Orchestrator.Tag,
		},
		&cli.StringFlag{
			Name:        "orchestrator-verbosity",
			Aliases:     []string{"orc-verbosity"},
			Usage:       "Sets orchestrator logging depth (Default: " + DefaultValues.Orchestrator.Verbosity + ")",
			Destination: &FlagValues.Orchestrator.Verbosity,
		},
		&cli.StringFlag{
			Name:        "orchestrator-vanguard-rpc-endpoint",
			Aliases:     []string{"orc-vanguard-rpc-endpoint"},
			Usage:       "Enables Vanguard node RPC provider endpoint.",
			Destination: &FlagValues.Orchestrator.VanguardRPCEndpoint,
		},
		&cli.StringFlag{
			Name:        "orchestrator-pandora-rpc-endpoint",
			Aliases:     []string{"orc-pandora-rpc-endpoint"},
			Usage:       "Pandora node RPC provider endpoint.",
			Destination: &FlagValues.Orchestrator.PandoraRPCEndpoint,
		},
		&cli.StringFlag{
			Name:        "pandora-tag",
			Aliases:     []string{"pan-tag", "pandora"},
			Usage:       "download and set pandora to given tag",
			Destination: &FlagValues.Pandora.Tag,
		},
		&cli.StringFlag{
			Name:        "pandora-verbosity",
			Aliases:     []string{"pan-verbosity"},
			Usage:       "Sets pandora logging depth (Default: " + DefaultValues.Pandora.Verbosity + ")",
			Destination: &FlagValues.Pandora.Verbosity,
		},
		&cli.StringFlag{
			Name:        "pandora-bootnodes",
			Aliases:     []string{"pan-bootnodes"},
			Usage:       "Sets pandora bootnodes (Default: " + DefaultValues.Pandora.Verbosity + ")",
			Destination: &FlagValues.Pandora.Bootnodes,
		},
		&cli.IntFlag{
			Name:        "pandora-port",
			Aliases:     []string{"pan-port"},
			Usage:       "Pandora client TCP/UDP port exposed. (Default:  " + strconv.Itoa(DefaultValues.Pandora.Port) + ")",
			Destination: &FlagValues.Pandora.Port,
		},
		&cli.StringFlag{
			Name:        "pandora-http-addr",
			Aliases:     []string{"pan-http-addr"},
			Usage:       "Pandora client http address exposed. (Default: " + DefaultValues.Pandora.HttpAddr + ")",
			Destination: &FlagValues.Pandora.Bootnodes,
		},
	}
	return luksoFlags
}
