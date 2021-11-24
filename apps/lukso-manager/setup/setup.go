package setup

import (
	"encoding/json"
	"fmt"
	"lukso-manager/downloader"
	"lukso-manager/shared"
	"net/http"
	"runtime"
	"strings"

	"github.com/boltdb/bolt"
)

type release struct {
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	URL    string `json:"url"`
	OsType string `json:"os_type"`
}

var ReleaseLocations = map[string]release{
	"pandora": {
		Name:   "Pandora",
		Tag:    "v0.2.0-rc.1",
		URL:    "https://github.com/lukso-network/pandora-execution-engine/releases/download/_TAG_/pandora-_OS_TYPE_",
		OsType: strings.Title(runtime.GOOS) + "-x86_64",
	},
	"vanguard": {
		Name:   "Vanguard",
		Tag:    "v0.2.0-rc.1",
		URL:    "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/_TAG_/vanguard-_OS_TYPE_",
		OsType: strings.Title(runtime.GOOS) + "-x86_64",
	},
	"lukso-orchestrator": {
		Name:   "Orchestrator",
		Tag:    "v0.2.0-rc.1",
		URL:    "https://github.com/lukso-network/lukso-orchestrator/releases/download/_TAG_/lukso-orchestrator-_OS_TYPE_",
		OsType: strings.Title(runtime.GOOS) + "-x86_64",
	},
	"lukso-deposit-cli": {
		Name:   "Deposit CLI",
		Tag:    "v1.2.6-LUKSO",
		URL:    "https://github.com/lukso-network/network-deposit-cli/releases/download/_TAG_/lukso-deposit-cli-_OS_TYPE_",
		OsType: strings.Title(runtime.GOOS) + "-x86_64",
	},
	"lukso-validator": {
		Name:   "Validator",
		Tag:    "v0.2.0-rc.1",
		URL:    "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/_TAG_/lukso-validator-_OS_TYPE_",
		OsType: strings.Title(runtime.GOOS) + "-x86_64",
	},
	"eth2stats": {
		Name:   "ETH 2 Stats",
		Tag:    "v0.2.0-rc.1",
		URL:    "https://github.com/lukso-network/network-vanguard-stats-client/releases/download/_TAG_/eth2stats-client-_OS_TYPE_",
		OsType: strings.Title(runtime.GOOS) + "-x86_64",
	},
}

type startClientsRequestBody struct {
	Network string
}

func Setup(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DOWNLOADING")
	decoder := json.NewDecoder(r.Body)
	var body startClientsRequestBody
	err := decoder.Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	downloader.DownloadConfigFiles(body.Network)
	for key, element := range ReleaseLocations {
		fmt.Println("Key:", key, "=>", "Element:", element)
		downloader.DownloadClientBinary(
			key,
			ReleaseLocations[key].Tag,
			ReleaseLocations[key].URL,
			ReleaseLocations[key].OsType,
		)
	}

	shared.SettingsDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("peers"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

}
