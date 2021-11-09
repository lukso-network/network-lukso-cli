package runner

import (
	"encoding/json"
	"log"
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

	errVanguard := startVanguard("v0.5.1-develop", network)
	if errVanguard != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errVanguard.Error()))
	}

	errOrchestrator := startOrchestrator("v0.5.4-develop", network)
	if errOrchestrator != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errOrchestrator.Error()))
	}

	errPandora := startPandora("v0.5.3-develop", network, body.Settings.HostName)
	if errPandora != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errPandora.Error()))
	}

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
