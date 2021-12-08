package pandora

import (
	"log"
	"lukso-cli/config"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

	err, NetworkConfig := config.LoadNetworkConfig(LuksoSettings.Network)
	if err != nil {
		log.Fatal("Config not loaded")
	}

	args := []string{
		"--datadir=" + LuksoSettings.DataDir + "/pandora",
		"--chainid=" + strconv.Itoa(NetworkConfig.ChainID),
		"--port=" + strconv.Itoa(LuksoSettings.Pandora.Port),
	}

	command := exec.Command("pandora" + strings.Join(args, " "))
	if startError := command.Start(); startError != nil {
		log.Fatal(startError)
		return
	}

	// command.Wait()
}

func Stop(LuksoSettings *config.LuksoValues) {
	println("Stopping Pandora")
}
