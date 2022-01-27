package main

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// TODO: consider extend it when new clients will be introduced
const (
	ELDependencyName            = "geth"
	ELGenesisDependencyName     = "genesis.json"
	CLDependencyName            = "prysm"
	validatorDependencyName     = "validator"
	CLGenesisDependencyName     = "genesis.ssz"
	CLConfigDependencyName      = "config.yml"
	CLStatsClientDependencyName = "eth2stats-client"
)

var (
	clientDependencies = map[string]*ClientDependency{
		ELDependencyName: {
			baseUrl: "https://github.com/lukso-network/go-ethereum/releases/download/%s/geth-%s-%s-%s",
			name:    ELDependencyName,
		},
		ELGenesisDependencyName: {
			baseUrl: "https://raw.githubusercontent.com/lukso-network/network-configs/l16-dev/l16/dev/4/genesis.json",
			name:    ELGenesisDependencyName,
		},
		CLDependencyName: {
			baseUrl: "https://github.com/lukso-network/prysm/releases/download/%s/beacon-chain-%s-%s-%s",
			name:    CLDependencyName,
		},
		validatorDependencyName: {
			baseUrl: "https://github.com/lukso-network/prysm/releases/download/%s/validator-%s-%s-%s",
			name:    validatorDependencyName,
		},
		CLGenesisDependencyName: {
			baseUrl: "https://raw.githubusercontent.com/lukso-network/network-configs/l16-dev/l16/dev/4/genesis.ssz",
			name:    CLGenesisDependencyName,
		},
		CLConfigDependencyName: {
			baseUrl: "https://raw.githubusercontent.com/lukso-network/network-configs/l16-dev/l16/dev/4/config.yaml",
			name:    CLConfigDependencyName,
		},
		CLStatsClientDependencyName: {
			baseUrl: "https://github.com/lukso-network/network-consensus-stats-client/releases/download/%s/eth2stats-client-%s-%s-%s",
			name:    CLStatsClientDependencyName,
		},
	}
)

type ClientDependency struct {
	baseUrl string
	name    string
}

func (dependency *ClientDependency) ParseUrl(tagName string) (url string) {
	sprintOccurrences := strings.Count(dependency.baseUrl, "%s")

	url = dependency.baseUrl
	if sprintOccurrences > 0 {
		url = fmt.Sprintf(dependency.baseUrl, tagName, tagName, systemOs, systemArch)
	}

	return
}

func (dependency *ClientDependency) ResolveDirPath(tagName string, datadir string) (location string) {
	location = fmt.Sprintf("%s/%s", datadir, tagName)

	return
}

func (dependency *ClientDependency) ResolveBinaryPath(tagName string, datadir string) (location string) {
	location = fmt.Sprintf("%s/%s", dependency.ResolveDirPath(tagName, datadir), dependency.name)

	return
}

func (dependency *ClientDependency) Run(
	tagName string,
	destination string,
	arguments []string,
	attachStdInAndErr bool,
) (pid int, err error) {
	binaryPath := dependency.ResolveBinaryPath(tagName, destination)
	command := exec.Command(binaryPath, arguments...)

	if attachStdInAndErr {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}

	err = command.Start()
	if err != nil {
		return 0, err
	}

	err = writePidToFile(destination, command.Process.Pid)
	if err != nil {
		return 0, err
	}

	return command.Process.Pid, nil
}

func (dependency *ClientDependency) Stop(destination string) (err error) {
	pid, err := getPidFromFile(destination)
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return
	}

	if pid == 0 {
		return nil
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return
	}

	err = process.Signal(syscall.SIGINT)
	if err != nil && !strings.Contains(err.Error(), "process already finished") {
		return
	}

	err = removePidFile(destination)
	if err != nil {
		return
	}

	return
}

func (dependency *ClientDependency) Download(tagName string, destination string) (err error) {
	dependencyTagPath := dependency.ResolveDirPath(tagName, destination)
	err = os.MkdirAll(dependencyTagPath, 0755)

	if nil != err {
		return
	}

	dependencyLocation := dependency.ResolveBinaryPath(tagName, destination)

	if fileExists(dependencyLocation) {
		log.Warningf("Not downloading dependency: %s, file already exists", dependencyLocation)

		return
	}

	fileUrl := dependency.ParseUrl(tagName)
	response, err := http.Get(fileUrl)

	if nil != err {
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf(
			"invalid response when downloading el on file url: %s. Response: %s",
			fileUrl,
			response.Status,
		)
	}

	output, err := os.Create(dependencyLocation)

	if nil != err {
		return
	}

	defer func() {
		_ = output.Close()
	}()

	_, err = io.Copy(output, response.Body)

	if nil != err {
		return
	}

	err = os.Chmod(dependencyLocation, os.ModePerm)

	return
}

func writePidToFile(path string, pid int) (err error) {
	fullFilepath := filepath.Join(path, PidFilename)
	s := big.NewInt(int64(pid))
	b := s.Bytes()
	err = os.WriteFile(fullFilepath, b, 0644)
	if err != nil {
		return
	}

	log.WithField("filepath", fullFilepath).Info("PID file written successfully")

	return
}

func getPidFromFile(path string) (pid int, err error) {
	fullFilepath := filepath.Join(path, PidFilename)
	b, err := os.ReadFile(fullFilepath)
	if err != nil {
		return 0, err
	}

	r := big.NewInt(0).SetBytes(b) // bytes to big Int
	pid = int(r.Int64())

	return pid, nil
}

func removePidFile(path string) (err error) {
	fullFilepath := filepath.Join(path, PidFilename)
	err = os.Remove(fullFilepath)
	if err != nil {
		return
	}

	return
}
