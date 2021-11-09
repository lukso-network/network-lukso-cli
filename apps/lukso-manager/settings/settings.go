package settings

import (
	"encoding/json"
	"fmt"
	"log"
	"lukso/shared"
	"net/http"

	"github.com/boltdb/bolt"
)

type Settings struct {
	HostName string `json:"hostName"`
	Coinbase string `json:"coinbase"`
}

type saveSettingsRequestBody struct {
	Network  string
	Settings Settings
}

func SaveSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body saveSettingsRequestBody
	errJson := decoder.Decode(&body)
	fmt.Println("body")
	fmt.Println(body.Settings.HostName)
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

	log.Println("Url Param 'network' is: " + string(network))

	settings, err := getSettings(shared.SettingsDB, network)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("returned settings")
	fmt.Println(settings)

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

func getSettings(db *bolt.DB, network string) (*Settings, error) {
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
