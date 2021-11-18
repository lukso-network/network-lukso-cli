package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"lukso/shared"
	"os"
	"path/filepath"
	"strings"
)

type DepositData []struct {
	Pubkey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Amount                int64  `json:"amount"`
	Signature             string `json:"signature"`
	DepositMessageRoot    string `json:"deposit_message_root"`
	DepositDataRoot       string `json:"deposit_data_root"`
	ForkVersion           string `json:"fork_version"`
	Eth2NetworkName       string `json:"eth2_network_name"`
	DepositCliVersion     string `json:"deposit_cli_version"`
}

func ReadDepositData(network string) DepositData {
	var depositData DepositData = nil

	validator_keys := shared.NetworkDir + network + "/validator_keys"

	err := filepath.Walk(validator_keys,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				if strings.Contains(path, "deposit_data") {
					depositData = readDepositJsonFile(path)

					if err != nil {
						log.Fatal(err)
					}
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return depositData
}

func readDepositJsonFile(path string) DepositData {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var depositData DepositData
	json.Unmarshal(byteValue, &depositData)

	return depositData
}
