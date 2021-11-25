package setup

import (
	"encoding/json"
	"fmt"
	"lukso/apps/lukso-manager/downloader"
	"lukso/apps/lukso-manager/shared"
	"net/http"

	"github.com/boltdb/bolt"
)

type release struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
	URL  string `json:"url"`
}

var defaultTag = "v0.1.0-develop"
var ReleaseLocations = map[string]release{
	"pandora": {
		Name: "Pandora",
		Tag:  defaultTag,
		URL:  "https://github.com/lukso-network/pandora-execution-engine/releases/download/_TAG_/pandora-_OS_TYPE_",
	},
	"vanguard": {
		Name: "Vanguard",
		Tag:  defaultTag,
		URL:  "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/_TAG_/vanguard-_OS_TYPE_",
	},
	"lukso-orchestrator": {
		Name: "Orchestrator",
		Tag:  defaultTag,
		URL:  "https://github.com/lukso-network/lukso-orchestrator/releases/download/_TAG_/lukso-orchestrator-_OS_TYPE_",
	},
	"lukso-deposit-cli": {
		Name: "Deposit CLI",
		Tag:  "v1.2.6-LUKSO",
		URL:  "https://github.com/lukso-network/network-deposit-cli/releases/download/_TAG_/lukso-deposit-cli-_OS_TYPE_",
	},
	"lukso-validator": {
		Name: "Validator",
		Tag:  defaultTag,
		URL:  "https://github.com/lukso-network/vanguard-consensus-engine/releases/download/_TAG_/lukso-validator-_OS_TYPE_",
	},
	"eth2stats": {
		Name: "ETH 2 Stats",
		Tag:  defaultTag,
		URL:  "https://github.com/lukso-network/network-vanguard-stats-client/releases/download/_TAG_/eth2stats-client-_OS_TYPE_",
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
		shared.HandleError(err, w)
		return
	}

	downloader.DownloadConfigFiles(body.Network)
	for key := range ReleaseLocations {
		downloader.DownloadClientBinary(
			key,
			ReleaseLocations[key].Tag,
			ReleaseLocations[key].URL,
		)
	}

	dbError := shared.SettingsDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("peers"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if dbError != nil {
		shared.HandleError(dbError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
