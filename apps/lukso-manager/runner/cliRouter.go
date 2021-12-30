package runner

import (
	"fmt"
	"log"
	"lukso/apps/lukso-manager/downloader"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/shared"
	"os"
)

func Prepare() error {

	LuksoSettings, err := settings.GetSettings(shared.SettingsDB, shared.PickedNetwork)
	if err != nil {
		log.Fatal(err)
	}

	// Download network configs if they are not present
	err = downloader.DownloadConfigFiles(shared.PickedNetwork)

	if err != nil {
		log.Fatal(err)
	}

	// Check whether binary is available and download if not present
	if LuksoSettings.Versions[settings.Orchestrator] != "" {
		// downloader.DownloadClient(string(settings.Orchestrator), LuksoSettings.Versions[settings.Orchestrator])
		// downloader.DownloadClient(string(settings.Orchestrator), "v0.1.0-rc.1-kravitz")
	}

	return nil
}

func HandleCli(cmd string, arg string) {

	luksoSettings, err := settings.GetSettings(shared.SettingsDB, shared.PickedNetwork)
	println(luksoSettings.Coinbase)
	println(luksoSettings.Versions[settings.Pandora])

	if err != nil {
		log.Fatal(err.Error())
	}

	err = Prepare()

	// downloader.GetAvailableVersionsDedicated()

	networkConfig, err := ReadConfig(shared.PickedNetwork)
	println(networkConfig.GENESISTIME)

	switch cmd {

	case "version":
		fmt.Println(shared.ScriptVersion)
		os.Exit(0)

	case "start":
		switch arg {
		case string(settings.Pandora):
			startPandora(luksoSettings.Versions[settings.Pandora], shared.PickedNetwork, *luksoSettings, networkConfig, fmt.Sprint(shared.RunningTime))
		}

	case "stop":
		switch arg {
		case string(settings.Pandora):
			stopPandora()
		}
	}
}
