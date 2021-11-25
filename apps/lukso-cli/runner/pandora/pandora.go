package pandora

import (
	"log"
	"lukso-cli/config"
	"os"
	"os/exec"
)

func Prepare(LuksoSettings *config.LuksoValues) {

	if _, err := os.Stat(LuksoSettings.LogsDir); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
		} else {
			// other error
		}
	}

}

func Start(LuksoSettings *config.LuksoValues) {
	Prepare(LuksoSettings)
	println("Starting Pandora")
	command := exec.Command("pandora")
	if startError := command.Start(); startError != nil {
		log.Fatal(startError)
		return
	}

	// command.Wait()
}

func Stop(LuksoSettings *config.LuksoValues) {
	println("Stopping Pandora")
}
