package cli

import (
	"log"
	"lukso/apps/lukso-manager/config"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/shared"

	"github.com/urfave/cli/v2"
)

func LoadFlags(c *cli.Context) {

	if c.Bool("api") {
		shared.EnableAPI = true
	}

	if c.Bool("gui") {
		shared.EnableGUI = true
	}

	shared.PickedNetwork = "l15-prod"

	networksNum := 0

	if flag := c.String("network"); flag != "" {
		shared.PickedNetwork = flag
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

	if flag := c.String("config"); flag != "" {
		config.LoadNodeConfig(LuksoSettings, flag)
	}

	if flag := c.String("coinbase"); flag != "" {
		LuksoSettings.Coinbase = flag
	}

	if flag := c.Bool("validate"); flag {
		LuksoSettings.ValidatorEnabled = flag
	}

	if flag := c.String("orchestrator"); flag != "" {
		LuksoSettings.Versions[settings.Orchestrator] = flag
	}

	if flag := c.String("pandora"); flag != "" {
		LuksoSettings.Versions[settings.Pandora] = flag
	}

	if flag := c.String("pandora-verbosity"); flag != "" {
		LuksoSettings.Pandora.Verbosity = flag
	}

	if flag := c.String("vanguard"); flag != "" {
		LuksoSettings.Versions[settings.Vanguard] = flag
	}

	if flag := c.String("validator"); flag != "" {
		LuksoSettings.Versions[settings.Validator] = flag
	}

	err2 := settings.SaveSettings(shared.SettingsDB, LuksoSettings, shared.PickedNetwork)

	if err2 != nil {
		log.Fatal(err2)
	}

}
