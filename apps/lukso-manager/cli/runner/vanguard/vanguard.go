package vanguard

import (
	"fmt"
	"lukso-cli/config"
	"os"
)

func prepare(LuksoSettings *config.LuksoValues) error {

	println(LuksoSettings.Network)

	err := os.MkdirAll(LuksoSettings.DataDir+"/vanguard", 0755)

	if err != nil {
		println(err.Error())
		return err
	}

	return nil

}

func Start(LuksoSettings *config.LuksoValues) {

	err := prepare(LuksoSettings)

	if err != nil {
		println(err)
		fmt.Println("Cannot start vanguard")
		// os.Exit(1)
	}

	println("Starting Vanguard...")

}

func Stop(LuksoSettings *config.LuksoValues) {
	println("Stopping Pandora")
}
