package settings

import (
	"encoding/json"
	"fmt"
	"log"
	"lukso/apps/lukso-manager/shared"
	"net/http"

	"github.com/boltdb/bolt"
)

type Client string

const (
	Vanguard     Client = "vanguard"
	Pandora      Client = "pandora"
	Orchestrator Client = "orchestrator"
	Validator    Client = "validator"
)

type Settings struct {
	HostName         string            `json:"hostName" yaml:"HOSTNAME"`
	Coinbase         string            `json:"coinbase" yaml:"COINBASE"`
	ExternalIP       string            `json:"externalIp" yaml:"EXTERNAL_IP"`
	Versions         map[Client]string `json:"versions" yaml:"VERSIONS"`
	ValidatorEnabled bool              `json:"validatorEnabled" yaml:"IS_VALIDATOR"`
}

type saveSettingsRequestBody struct {
	Network  string
	Settings Settings
}

func SaveSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body saveSettingsRequestBody

	if body.Settings.ExternalIP == "" {
		body.Settings.ExternalIP = shared.OutboundIP.String()
	}

	errJson := decoder.Decode(&body)
	if errJson != nil {
		panic(errJson)
	}

	err := SaveSettings(shared.SettingsDB, &body.Settings, body.Network)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode("Settings successfuly saved"); err != nil {
		panic(err)
	}
}

func SaveSettings(db *bolt.DB, settings *Settings, network string) error {
	// Store the user model in the user bucket using the username as the key.
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(network))
		if err != nil {
			fmt.Println(err)
			return err
		}

		encoded, err := json.Marshal(settings)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return b.Put([]byte("settings"), encoded)
	})

	return err
}

func GetSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["network"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	network := keys[0]

	settings, err := GetSettings(shared.SettingsDB, network)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(settings); err != nil {
		panic(err)
	}
}

func decodeSettings(data []byte) (*Settings, error) {
	var settings *Settings
	err := json.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func GetSettings(db *bolt.DB, network string) (*Settings, error) {
	// Store the user model in the user bucket using the username as the key.
	var settings *Settings
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(network))
		k := []byte("settings")
		settings, err = decodeSettings(b.Get(k))

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Could not get settings")
		return nil, err
	}
	return settings, nil

}

func DefaultSettings(db *bolt.DB, network string) error {
    var settings Settings
	settings.Coinbase = ""
	settings.ValidatorEnabled = false
	settings.ExternalIP = ""
	settings.HostName = ""
	settings.Versions = map[Client]string{
		Orchestrator: "v0.2.0-rc.2",
		Pandora: "v0.2.0-rc.2",
		Vanguard: "v0.2.0-rc.2",
		Validator: "v0.2.0-rc.2",
	}
	SaveSettings(db, &settings, network)
	return nil
}