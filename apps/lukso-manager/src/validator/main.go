package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"lukso/apps/lukso-manager/src/runner"
	"net/http"
	"os"

	"github.com/tyler-smith/go-bip39"
)

type generateValidatorKeysRequestBody struct {
	Password           string
	Network            string
	AmountOfValidators string
}

func GenerateValidatorKeys(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body generateValidatorKeysRequestBody
	errJson := decoder.Decode(&body)
	if errJson != nil {
		panic(errJson)
	}

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	// seed := bip39.NewSeed(mnemonic, password)

	folder := "/home/rryter/.lukso/networks/" + body.Network + "/validator_keys"

	err := os.Chmod(folder, 0775)
	if err != nil {
		os.Mkdir(folder, 0775)
	}

	mnemonicData := []byte(mnemonic)
	errWrite := os.WriteFile(folder+"/mnemonic", mnemonicData, 0644)
	if errWrite != nil {
		log.Fatal("write failed")
	}

	args := []string{
		"existing-mnemonic",
		"--num_validators " + body.AmountOfValidators,
		"--folder '" + folder + "'",
		"--keystore_password " + body.Password,
		"--validator_start_index 0",
		"--chain " + body.Network,
		"--mnemonic '" + mnemonic + "'",
	}

	fmt.Println(body.AmountOfValidators)
	fmt.Println(args)

	runner.StartBinary("lukso-deposit-cli", "v1.2.6-LUKSO", args)
}
