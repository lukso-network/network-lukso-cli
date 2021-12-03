package vanguard

import (
	"fmt"
	"lukso-cli/config"
	"os"
)

func prepare(LuksoSettings *config.LuksoValues) error {

	err := os.Mkdir(LuksoSettings.Network+"/vanguard", 0755)
	println(LuksoSettings.DataDir)

	if err != nil {
		return fmt.Errorf("")
	}

	return nil

}

func Start(LuksoSettings *config.LuksoValues) {

	err := prepare(&*LuksoSettings)

	if err != nil {
		fmt.Println("Cannot start vanguard")
		// os.Exit(1)
	}

	println("Starting Vanguard...")

}

func Stop(LuksoSettings *config.LuksoValues) {
	println("Stopping Pandora")
}
