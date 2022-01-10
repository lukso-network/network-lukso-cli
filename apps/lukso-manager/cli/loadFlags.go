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

	if c.Bool("validate") {
		LuksoSettings.ValidatorEnabled = true
	}

	if c.String("orchestrator") != "" {
		LuksoSettings.Versions[settings.Pandora] = c.String("orchestrator")
	}

	if c.String("pandora") != "" {
		LuksoSettings.Versions[settings.Pandora] = c.String("pandora")
	}

	if c.String("vanguard") != "" {
		LuksoSettings.Versions[settings.Vanguard] = c.String("vanguard")
	}

	if c.String("validator") != "" {
		LuksoSettings.Versions[settings.Vanguard] = c.String("validator")
	}

	err2 := settings.SaveSettings(shared.SettingsDB, LuksoSettings, shared.PickedNetwork)

	if err2 != nil {
		log.Fatal(err2)
	}

}
