package runner

import (
	"encoding/json"
	"log"
	"lukso-manager/downloader"
	"lukso-manager/settings"
	"lukso-manager/shared"
	"net/http"
	"os"
	"os/exec"
)

type startClientsRequestBody struct {
	Network  string
	Clients  []string
	Settings settings.Settings
}

type stopClientsRequestBody struct {
	Clients []string
}

type Commands struct {
	orchestrator *exec.Cmd
	pandora      *exec.Cmd
	vanguard     *exec.Cmd
	validator    *exec.Cmd
}

var CommandsByClient = Commands{}

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

	oldConfig, _ := ReadConfig(network)
	downloader.DownloadConfigFiles(network)
	newConfig, _ := ReadConfig(network)

	if oldConfig.GENESISTIME != newConfig.GENESISTIME {
		err := os.RemoveAll(shared.NetworkDir + network + "/" + shared.DataDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	if shared.Contains(body.Clients, "vanguard") {
		vanCmd, errVanguard := startVanguard(body.Settings.Versions[settings.Vanguard], network)
		if errVanguard != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errVanguard.Error()))
			return
		}
		CommandsByClient.vanguard = vanCmd
	}

	if shared.Contains(body.Clients, "orchestrator") {
		orchCmd, errOrchestrator := startOrchestrator(body.Settings.Versions[settings.Orchestrator], network)
		if errOrchestrator != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errOrchestrator.Error()))
			return
		}
		CommandsByClient.orchestrator = orchCmd
	}

	if shared.Contains(body.Clients, "pandora") {
		cmdPandora, errPandora := startPandora(body.Settings.Versions[settings.Pandora], network, body.Settings)
		if errPandora != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errPandora.Error()))
			return
		}
		CommandsByClient.pandora = cmdPandora
	}

	if shared.Contains(body.Clients, "validator") {
		cmdValidator, errValidator := startValidator(body.Settings.Versions[settings.Validator], network)
		if errValidator != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errValidator.Error()))
			return
		}
		CommandsByClient.validator = cmdValidator
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully started all the clients."); err != nil {
		panic(err)
	}
}

func StopClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body stopClientsRequestBody
	err := decoder.Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if shared.Contains(body.Clients, "pandora") && CommandsByClient.pandora != nil {
		if err := CommandsByClient.pandora.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	if shared.Contains(body.Clients, "vanguard") {
		if err := CommandsByClient.vanguard.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	if shared.Contains(body.Clients, "orchestrator") {
		if err := CommandsByClient.orchestrator.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	if shared.Contains(body.Clients, "validator") {
		if err := CommandsByClient.validator.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func StartBinary(client string, version string, args []string) (*exec.Cmd, error) {

	command := exec.Command(shared.BinaryDir+client+"/"+version+"/"+client, args...)

	if startError := command.Start(); startError != nil {
		log.Println("ERROR STARTING " + client + "@" + version)
		log.Fatal(startError)
		return nil, startError
	}

	return command, nil

}
