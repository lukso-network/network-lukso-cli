package metrics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lukso/shared"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
)

func VanguardMetrics(w http.ResponseWriter, r *http.Request) {
	body, err := getMetrics("http://127.0.0.1:8080/metrics", w)
	fmt.Println("body")

	handleError(err, w)
	returnBody(body, w)
}

func PandoraMetrics(w http.ResponseWriter, r *http.Request) {
	body, err := getMetrics("http://127.0.0.1:6060/debug/metrics", w)
	var myStoredVariable map[string]int64

	json.Unmarshal(body, &myStoredVariable)

	err2 := shared.SettingsDB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("peers"))
		if err != nil {
			fmt.Println(err)
			return err
		}

		peersOverTime, _ := getPeersOverTime()
		if peersOverTime == nil {
			peersOverTime = make(map[int64]int64)
		}

		now := time.Now()
		sec := now.Unix()

		peersOverTime[sec] = myStoredVariable["p2p/peers"]

		a, _ := json.Marshal(peersOverTime)

		return b.Put([]byte("pandoraPeers"), a)
	})

	handleError(err, w)
	handleError(err2, w)

	returnBody(body, w)
}

func ValidatorMetrics(w http.ResponseWriter, r *http.Request) {
	body, err := getMetrics("http://127.0.0.1:8081/metrics", w)
	handleError(err, w)
	returnBody(body, w)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func getMetrics(url string, w http.ResponseWriter) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return
	}

	return
}

func GetPandoraPeersOverTime(w http.ResponseWriter, r *http.Request) {
	metrics, err := getPeersOverTime()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonString, _ := json.Marshal(metrics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func returnBody(body []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func decodeSettings(data []byte) (metrics map[int64]int64, err error) {
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		return
	}
	return
}

func getPeersOverTime() (map[int64]int64, error) {
	// Store the user model in the user bucket using the username as the key.
	var settings map[int64]int64
	err := shared.SettingsDB.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("peers"))
		k := []byte("pandoraPeers")
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
