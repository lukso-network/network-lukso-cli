package cli

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"lukso/apps/lukso-manager/config"
	"lukso/apps/lukso-manager/runner"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/shared"
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

	if c.Bool("api") {
		shared.EnableAPI = true
	}

	if c.Bool("gui") {
		shared.EnableGUI = true
	}

	shared.PickedNetwork = "l15-prod"

	networksNum := 0

	if c.String("network") != "" {
		shared.PickedNetwork = c.String("network")
		networksNum++
	}

	if c.Bool("l15-prod") {
		shared.PickedNetwork = "l15-prod"
		networksNum++
	}

	if c.Bool("l15-staging") {
		shared.PickedNetwork = "l15-staging"
		networksNum++
	}

	if c.Bool("l15-dev") {
		shared.PickedNetwork = "l15-dev"
		networksNum++
	}

	if networksNum > 1 {
		log.Fatal("ERROR: You cannot connect to multiple networks, please choose only one.")
	}

	println(shared.PickedNetwork)

	err := settings.DefaultSettings(shared.SettingsDB, shared.PickedNetwork)

	if err != nil {
		log.Fatal(err)
	}

	LuksoSettings, err0 := settings.GetSettings(shared.SettingsDB, shared.PickedNetwork)

	if err0 != nil {
		log.Fatal(err0)
	}

	if c.String("config") != "" {
		config.LoadNodeConfig(LuksoSettings, c.String("config"))
	}

	if c.String("coinbase") != "" {
		LuksoSettings.Coinbase = c.String("coinbase")
	}

	err2 := settings.SaveSettings(shared.SettingsDB, LuksoSettings, shared.PickedNetwork)

	if err2 != nil {
		log.Fatal(err2)
	}

}
