package cli

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"lukso/apps/lukso-manager/runner"
	"lukso/apps/lukso-manager/settings"
)

var Cmd string
var Arg string

var API bool
var GUI bool

func Init() {

	app := cli.NewApp()
	app.Name = "LUKSO CLI"
	app.Usage = "Tool for managing LUKSO node"
	app.UsageText = "lukso <command> [argument] [--flags]"
	app.Flags = getLuksoFlags()
	app.EnableBashCompletion = true
	app.After = func(c *cli.Context) error {
		LoadFlags(c)
		return nil
	}

	app.Commands = []*cli.Command{
		getStartCommand(),
		getStopCommand(),
		getVersionCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	runner.HandleCli(Cmd, Arg)

}

// Loads flag values into settings struct
func LoadFlags(c *cli.Context) {

	var network string

	networksNum := 0

	if c.Bool("GUI") {
		GUI = true
	}

	if c.String("network") != "" {
		network = c.String("network")
		networksNum++
	}

	if c.Bool("l15-prod") {
		network = "l15-prod"
		networksNum++
	}

	if c.Bool("l15-staging") {
		network = "l15-staging"
		networksNum++
	}

	if c.Bool("l15-dev") {
		network = "l15-dev"
		networksNum++
	}

	if networksNum > 1 {
		log.Fatal("ERROR: You cannot connect to multiple networks, please choose only one.")
	}

	var LuksoSettings settings.Settings
	println(network)

	if c.String("coinbase") != "" {
		LuksoSettings.Coinbase = c.String("coinbase")
	}

}
