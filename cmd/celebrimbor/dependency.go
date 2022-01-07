package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// TODO: consider to move it to common/shared
const (
	ELDependencyName        = "el"
	ELGenesisDependencyName = "el_private_testnet_genesis.json"
	CLDependencyName        = "vanguard"
	validatorDependencyName = "validator"
	CLGenesisDependencyName = "vanguard_private_testnet_genesis.ssz"
	CLConfigDependencyName  = "config.yml"
)

var (
	clientDependencies = map[string]*ClientDependency{
		ELDependencyName: {
			baseUnixUrl:   "https://github.com/lukso-network/pandora-execution-engine/releases/download/%s/geth",
			baseDarwinUrl: "https://github.com/lukso-network/pandora-execution-engine/releases/download/%s/geth-darwin",
			name:          ELDependencyName,
		},
		ELGenesisDependencyName: {
			baseUnixUrl:   "https://storage.googleapis.com/l16-common/pandora/pandora_private_testnet_genesis.json",
			baseDarwinUrl: "https://storage.googleapis.com/l16-common/pandora/pandora_private_testnet_genesis.json",
			name:          ELGenesisDependencyName,
		},
		CLDependencyName: {
			baseUnixUrl:   "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/%s/beacon-chain",
			baseDarwinUrl: "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/%s/beacon-chain-darwin",
			name:          CLDependencyName,
		},
		validatorDependencyName: {
			baseUnixUrl:   "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/%s/validator",
			baseDarwinUrl: "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/%s/validator-darwin",
			name:          validatorDependencyName,
		},
		CLGenesisDependencyName: {
			baseUnixUrl:   "https://storage.googleapis.com/l16-common/vanguard/vanguard_private_testnet_genesis.ssz",
			baseDarwinUrl: "https://storage.googleapis.com/l16-common/vanguard/vanguard_private_testnet_genesis.ssz",
			name:          CLGenesisDependencyName,
		},
		CLConfigDependencyName: {
			baseUnixUrl:   "https://storage.googleapis.com/l16-common/vanguard/chain-config.yaml",
			baseDarwinUrl: "https://storage.googleapis.com/l16-common/vanguard/chain-config.yaml",
			name:          CLConfigDependencyName,
		},
	}
)

type ClientDependency struct {
	baseUnixUrl   string
	baseDarwinUrl string
	name          string
}

func (dependency *ClientDependency) ParseUrl(tagName string) (url string) {
	// do not parse when no occurrences
	sprintOccurrences := strings.Count(dependency.baseUnixUrl, "%s")
	currentOs := systemOs

	if sprintOccurrences < 1 && currentOs == ubuntu {
		return dependency.baseUnixUrl
	}

	if sprintOccurrences < 1 && currentOs == macos {
		return dependency.baseDarwinUrl
	}

	if currentOs == macos {
		return fmt.Sprintf(dependency.baseDarwinUrl, tagName)
	}

	return fmt.Sprintf(dependency.baseUnixUrl, tagName)
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
) (err error) {
	binaryPath := dependency.ResolveBinaryPath(tagName, destination)
	command := exec.Command(binaryPath, arguments...)

	if attachStdInAndErr {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}

	err = command.Start()

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
		log.Warning("I am not downloading el, file already exists")

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
