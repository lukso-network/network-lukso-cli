package metrics

import (
	"encoding/json"
	"fmt"
	"lukso/apps/lukso-manager/shared"
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

func setPeersOverTime(amountOfPeers float64, client string) (err error) {
	errDbUpdate := shared.SettingsDB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("peers"))
		if err != nil {
			fmt.Println(err)
			return err
		}

		peersOverTime, _ := getPeersOverTime(client)
		if peersOverTime == nil {
			peersOverTime = make(map[int64]float64)
		}

		sortingInts := make([]float64, 0)

		for key := range peersOverTime {
			sortingInts = append(sortingInts, float64(key))
		}

		sort.Float64s(sortingInts)

		if len(sortingInts) > 100 {
			sortingInts = sortingInts[len(sortingInts)-100:]
		}

		reducedMap := make(map[int64]float64)

		for _, timestamp := range sortingInts {
			reducedMap[int64(timestamp)] = peersOverTime[int64(timestamp)]
		}

		now := time.Now()
		sec := now.Unix()

		reducedMap[sec] = amountOfPeers

		json, _ := json.Marshal(reducedMap)

		return b.Put([]byte(client+"peersOverTime"), json)
	})

	if errDbUpdate != nil {
		return
	}

	return
}

func getPeersOverTime(client string) (map[int64]float64, error) {
	// Store the user model in the user bucket using the username as the key.
	var peersOverTime map[int64]float64
	err := shared.SettingsDB.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("peers"))
		k := []byte(client + "peersOverTime")
		peersOverTime, err = decodeSettings(b.Get(k))

		if err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Could not get settings")
		return nil, err
	}
	return peersOverTime, nil

}
