package runner

import (
	"encoding/json"
	"log"
	"lukso/downloader"
	"lukso/settings"
	"lukso/shared"
	"net/http"
	"os/exec"
)

type startClientsRequestBody struct {
	Network  string
	Settings settings.Settings
}

func StartClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body startClientsRequestBody
	err := decoder.Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	network := body.Network

	downloader.DownloadConfigFiles(network)

	errVanguard := startVanguard(body.Settings.Versions[settings.Vanguard], network)
	if errVanguard != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errVanguard.Error()))
	}

	errOrchestrator := startOrchestrator(body.Settings.Versions[settings.Orchestrator], network)
	if errOrchestrator != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errOrchestrator.Error()))
	}

	errPandora := startPandora(body.Settings.Versions[settings.Pandora], network, body.Settings)
	if errPandora != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errPandora.Error()))
	}

	// if body.Settings.ValidatorEnabled {
	errValidator := startValidator(body.Settings.Versions[settings.Validator], network)
	if errValidator != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errValidator.Error()))
	}
	// }

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully started all the clients."); err != nil {
		panic(err)
	}
}

func StopClients(w http.ResponseWriter, r *http.Request) {
	command := exec.Command("lukso", "stop")

	if err := command.Start(); err != nil {
		log.Fatal(err)
	}
}

func StartBinary(client string, version string, args []string) {

	command := exec.Command(shared.BinaryDir+client+"/"+version+"/"+client, args...)

	if startError := command.Start(); startError != nil {
		log.Println("ERROR STARTING " + client + "@" + version)
		log.Fatal(startError)
	}

}
