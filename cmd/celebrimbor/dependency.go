package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// TODO: consider extend it when new clients will be introduced
const (
	ELDependencyName        = "geth"
	ELGenesisDependencyName = "el_private_testnet_genesis.json"
	CLDependencyName        = "prysm"
	validatorDependencyName = "validator"
	CLGenesisDependencyName = "cl_private_testnet_genesis.ssz"
	CLConfigDependencyName  = "config.yml"
)

var (
	clientDependencies = map[string]*ClientDependency{
		ELDependencyName: {
			baseUnixUrl:   "https://github.com/silesiacoin/go-ethereum/releases/download/%s/geth-Linux-x86_64",
			baseDarwinUrl: "https://github.com/silesiacoin/go-ethereum/releases/download/%s/geth-Darwin-x86_64",
			name:          ELDependencyName,
		},
		ELGenesisDependencyName: {
			baseUnixUrl:   "https://storage.googleapis.com/merge-network/configs/templates/genesis.json",
			baseDarwinUrl: "https://storage.googleapis.com/merge-network/configs/templates/genesis.json",
			name:          ELGenesisDependencyName,
		},
		CLDependencyName: {
			baseUnixUrl:   "https://github.com/silesiacoin/prysm/releases/download/%s/beacon-chain-Linux-x86_64",
			baseDarwinUrl: "https://github.com/silesiacoin/prysm/releases/download/%s/beacon-chain-Darwin-x86_64",
			name:          CLDependencyName,
		},
		validatorDependencyName: {
			baseUnixUrl:   "https://github.com/silesiacoin/prysm/releases/download/%s/validator-Darwin-x86_64",
			baseDarwinUrl: "https://github.com/silesiacoin/prysm/releases/download/%s/validator-Linux-x86_64",
			name:          validatorDependencyName,
		},
		CLGenesisDependencyName: {
			baseUnixUrl:   "https://storage.googleapis.com/merge-network/configs/templates/genesis.ssz",
			baseDarwinUrl: "https://storage.googleapis.com/merge-network/configs/templates/genesis.ssz",
			name:          CLGenesisDependencyName,
		},
		CLConfigDependencyName: {
			baseUnixUrl:   "https://storage.googleapis.com/merge-network/configs/templates/config.yaml",
			baseDarwinUrl: "https://storage.googleapis.com/merge-network/configs/templates/config.yaml",
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

	// Fallback to unix if system is not recognized
	if sprintOccurrences < 1 {
		return dependency.baseUnixUrl
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
