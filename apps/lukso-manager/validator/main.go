package validator

import (
	"bufio"
	"encoding/json"
	"log"
	"lukso/apps/lukso-manager/shared"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

type generateValidatorKeysRequestBody struct {
	Password           string
	Network            string
	AmountOfValidators string
}

type importValidatorKeysRequestBody struct {
	Network        string
	KeysPassword   string
	WalletPassword string
}

type resetValidatorKeysRequestBody struct {
	Network string
}

func ResetValidator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body resetValidatorKeysRequestBody
	errJson := decoder.Decode(&body)
	if errJson != nil {
		panic(errJson)
	}

	validator_keys := shared.NetworkDir + body.Network + "/validator_keys"
	_, errV := os.Stat(validator_keys)
	if errV == nil {
		zipFolder(body.Network, "validator_keys")
		os.RemoveAll(validator_keys)
	}

	vanguard_wallet := shared.NetworkDir + body.Network + "/vanguard_wallet"
	_, errW := os.Stat(vanguard_wallet)
	if errW == nil {
		zipFolder(body.Network, "vanguard_wallet")
		os.RemoveAll(vanguard_wallet)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully removed keys and wallet"); err != nil {
		panic(err)
	}
}

func GenerateValidatorKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body generateValidatorKeysRequestBody
	errJson := decoder.Decode(&body)
	if errJson != nil {
		shared.HandleError(errJson, w)
		return
	}

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	folder := shared.NetworkDir + body.Network

	_, err := os.Stat(folder)
	if err != nil {
		os.MkdirAll(folder, 0775)
	}

	_, err1 := os.Stat(folder + "/validator_keys")
	if err1 != nil {
		os.MkdirAll(folder, 0775)
	}

	_, errPw := os.Stat(folder + "/passwords")
	if errPw != nil {
		os.MkdirAll(folder+"/passwords", 0775)
	}

	mnemonicData := []byte(mnemonic)
	errWrite := os.WriteFile(folder+"/passwords/mnemonic", mnemonicData, 0644)
	if errWrite != nil {
		log.Fatal("write failed for " + folder + "/passwords/mnemonic")
		shared.HandleError(errWrite, w)
		return
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

	errPW := os.WriteFile(folder+"/passwords/keys", []byte(body.Password), 0644)
	if errPW != nil {
		shared.HandleError(errPW, w)
		return
	}

	command := exec.Command("bash", "-c", shared.BinaryDir+"lukso-deposit-cli/v1.2.6-LUKSO/lukso-deposit-cli "+strings.Join(args, " "))

	if startError := command.Start(); startError != nil {
		shared.HandleError(startError, w)
		return
	}

	command.Wait()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully created keys"); err != nil {
		shared.HandleError(err, w)
		return
	}
}

func ImportValidatorKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body importValidatorKeysRequestBody
	errJson := decoder.Decode(&body)
	if errJson != nil {
		shared.HandleError(errJson, w)
		return
	}

	folder := shared.NetworkDir + body.Network
	validatorKeysFolder := folder + "/validator_keys"
	walletFolder := folder + "/vanguard_wallet"
	passwordFolder := folder + "/passwords"

	args := []string{
		"accounts",
		"import",
		"--wallet-dir " + walletFolder,
		"--keys-dir " + validatorKeysFolder,
		"--wallet-password-file " + passwordFolder + "/keys",
		"--account-password-file " + passwordFolder + "/keys",
	}

	command := exec.Command("bash", "-c", shared.BinaryDir+"lukso-validator/v0.5.3-develop/lukso-validator "+strings.Join(args, " "))

	stdout, _ := command.StdoutPipe()

	if startError := command.Start(); startError != nil {
		shared.HandleError(startError, w)
		return
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Println(in.Text())
	}

	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}

	command.Wait()

	compressedValidatorKeys := zipFolder(body.Network, "validator_keys")

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("validator_keys.zip"))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, compressedValidatorKeys)
}

func GetDepositData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	keys, ok := r.URL.Query()["network"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	network := keys[0]

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ReadDepositData(network)); err != nil {
		panic(err)
	}
}
