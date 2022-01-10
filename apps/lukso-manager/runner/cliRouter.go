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

	downloadedVersions, err := downloader.GetDownloadedVersions()

	if err != nil {
		log.Fatal(err)
	}

	if orchestratorVersion := LuksoSettings.Versions[settings.Orchestrator]; !(shared.Contains(downloadedVersions[string(settings.Orchestrator)], orchestratorVersion)) {
		fmt.Println("Downloading:", "Orchestrator with tag", orchestratorVersion)
		downloader.DownloadClient(string(settings.Orchestrator), orchestratorVersion)
	}

	if pandoraVersion := LuksoSettings.Versions[settings.Pandora]; !(shared.Contains(downloadedVersions[string(settings.Pandora)], pandoraVersion)) {
		fmt.Println("Downloading:", "Pandora with tag", pandoraVersion)
		downloader.DownloadClient(string(settings.Pandora), pandoraVersion)
	}

	if vanguardVersion := LuksoSettings.Versions[settings.Vanguard]; !(shared.Contains(downloadedVersions[string(settings.Vanguard)], vanguardVersion)) {
		fmt.Println("Downloading:", "Vanguard with tag", vanguardVersion)
		downloader.DownloadClient(string(settings.Vanguard), vanguardVersion)
	}

	if validatorVersion := LuksoSettings.Versions[settings.Validator]; !(shared.Contains(downloadedVersions[string(settings.Validator)], validatorVersion)) {
		fmt.Println("Downloading:", "Validator with tag", validatorVersion)
		downloader.DownloadClient(string(settings.Validator), validatorVersion)
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

	networkConfig, err := ReadNetworkConfig(shared.PickedNetwork)

	switch cmd {

	case "version":
		fmt.Println(shared.ScriptVersion)
		os.Exit(0)

	case "start":
		switch arg {
		case string(settings.Orchestrator):
			startOrchestrator(luksoSettings.Versions[settings.Orchestrator], shared.PickedNetwork)
		case string(settings.Pandora):
			startPandora(luksoSettings.Versions[settings.Pandora], shared.PickedNetwork, *luksoSettings, networkConfig, fmt.Sprint(shared.RunningTime))
		case string(settings.Vanguard):
			startVanguard(luksoSettings.Versions[settings.Vanguard], shared.PickedNetwork, networkConfig, fmt.Sprint(shared.RunningTime))
		case string(settings.Validator):
			startValidator(luksoSettings.Versions[settings.Validator], shared.PickedNetwork, networkConfig, fmt.Sprint(shared.RunningTime))
		}

	case "stop":
		switch arg {
		case string(settings.Pandora):
			stopPandora()
		}
	}
}
