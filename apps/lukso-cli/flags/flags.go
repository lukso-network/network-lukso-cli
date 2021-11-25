package flags

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"lukso-cli/config"
)

var Cmd string
var Arg string

func InitFlags() {

	app := cli.NewApp()
	app.Name = "LUKSO CLI"
	app.Usage = "Tool for managing LUKSO node"
	app.UsageText = "lukso <command> [argument] [--flags]"
	app.Flags = getLuksoFlags()
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		getStartCommand(),
		getStopCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func LoadFlags(LuksoSettings *config.LuksoValues) {

	networksNum := 0

	if FlagValues.Network != "" {
		networksNum++
	}

	if FlagValues.l15_prod {
		networksNum++
	}

	if FlagValues.l15_staging {
		networksNum++
	}

	if FlagValues.l15_dev {
		networksNum++
	}

	if networksNum > 1 {
		log.Fatal("ERROR: You cannot connect to multiple networks, please choose only one.")
	}

	if FlagValues.Network != "" {
		LuksoSettings.Network = FlagValues.Network
	}

	if FlagValues.l15_prod {
		LuksoSettings.Network = "l15-prod"
	}

	if FlagValues.l15_staging {
		LuksoSettings.Network = "l15-staging"
	}

	if FlagValues.l15_dev {
		LuksoSettings.Network = "l15-dev"
	}

	if FlagValues.Orchestrator.Verbosity != "" {
		LuksoSettings.Orchestrator.Verbosity = FlagValues.Orchestrator.Verbosity
	}

}
