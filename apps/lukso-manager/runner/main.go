package runner

import (
	"encoding/json"
	"fmt"
	"log"
	"lukso-manager/downloader"
	"lukso-manager/settings"
	"lukso-manager/shared"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
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
	config, _ := ReadConfig(network)

	if oldConfig.GENESISTIME != config.GENESISTIME {
		err := os.RemoveAll(shared.NetworkDir + network + "/" + shared.DataDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	folder := shared.NetworkDir + body.Network + "/logs/" + fmt.Sprint(config.GENESISTIME)

	_, folderErr := os.Stat(folder)
	if folderErr != nil {
		os.Mkdir(folder, 0775)
	}

	now := time.Now()
	timestamp := now.Unix()

	if shared.Contains(body.Clients, "vanguard") {
		vanCmd, errVanguard := startVanguard(body.Settings.Versions[settings.Vanguard], network, config, fmt.Sprint(timestamp))
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
		version := body.Settings.Versions[settings.Pandora]
		cmdPandora, errPandora := startPandora(version, network, body.Settings, config, fmt.Sprint(timestamp))
		if errPandora != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errPandora.Error()))
			return
		}
		CommandsByClient.pandora = cmdPandora
	}

	if shared.Contains(body.Clients, "validator") {
		version := body.Settings.Versions[settings.Validator]
		cmdValidator, errValidator := startValidator(version, network, config, fmt.Sprint(timestamp))
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

	log.Println("STARTING " + client + "@" + version)
	log.Println("ARGS " + strings.Join(args, " "))
	command := exec.Command(shared.BinaryDir+client+"/"+version+"/"+client, args...)

	if startError := command.Start(); startError != nil {
		log.Println("ERROR STARTING " + client + "@" + version)
		log.Fatal(startError)
		return nil, startError
	}

	return command, nil

}
