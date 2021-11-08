package validator

import (
	"encoding/json"
	"log"
	"lukso/runner"
	"lukso/shared"
	"net/http"
	"os"
	"os/exec"
	"strings"

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

	folder := shared.NetworkDir + body.Network

	err := os.Chmod(folder, 0775)
	if err != nil {
		os.Mkdir(folder, 0775)
	}

	mnemonicData := []byte(mnemonic)
	errWrite := os.WriteFile(folder+"/validator_keys/mnemonic", mnemonicData, 0644)
	if errWrite != nil {
		log.Fatal("write failed")
	}

	args := []string{
		"existing-mnemonic",
		"--folder '" + folder + "'",
		"--num_validators " + body.AmountOfValidators,
		"--keystore_password " + body.Password,
		"--validator_start_index 0",
		"--chain " + body.Network,
		"--mnemonic '" + mnemonic + "'",
	}

	runner.StartBinary("lukso-deposit-cli", "v1.2.6-LUKSO", args)

	command := exec.Command("bash", "-c", shared.BinaryDir+"lukso-deposit-cli/v1.2.6-LUKSO/lukso-deposit-cli "+strings.Join(args, " "))

	if startError := command.Start(); startError != nil {
		log.Fatal(startError)
	}

	command.Wait()
}
