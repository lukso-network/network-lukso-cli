package runner

import (
	"encoding/json"
	"log"
	"lukso/shared"
	"net/http"
	"os/exec"
)

func StartClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	network := "l15-staging"

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

	errPandora := startPandora("v0.5.3-develop", network)
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
