package runner

import (
	"fmt"
	"log"
	"lukso/apps/lukso-manager/downloader"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/shared"
	"os"
)

func HandleCli(cmd string, arg string) {

	luksoSettings, err := settings.GetSettings(shared.SettingsDB, shared.PickedNetwork)
	println(luksoSettings.Coinbase)
	println(luksoSettings.Versions[settings.Pandora])

	if err != nil {

	}

	err = downloader.DownloadConfigFiles(shared.PickedNetwork)

	if err != nil {
		log.Fatal(err)
	}
	networkConfig, err := ReadConfig(shared.PickedNetwork)

	if err != nil {
		log.Fatal(err)
	}

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
