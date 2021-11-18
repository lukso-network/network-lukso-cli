package validator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"lukso/shared"
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

func ResetValidator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body generateValidatorKeysRequestBody
	errJson := decoder.Decode(&body)
	if errJson != nil {
		panic(errJson)
	}

	folder := shared.NetworkDir + body.Network

	zipFolder(body.Network, "validator_keys")
	errDeleteValidatoKeys := os.Remove(folder + "/validator_keys")
	if errDeleteValidatoKeys != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode("Could not delete validator keys"); err != nil {
			panic(err)
		}
		return
	}

	zipFolder(body.Network, "vanguard_wallet")
	errDeleteVanguardWallet := os.Remove(folder + "/vanguard_wallet")
	if errDeleteVanguardWallet != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode("Could not delete validator wallet"); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully created keys"); err != nil {
		panic(err)
	}
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

	_, err := os.Stat(folder)
	if err != nil {
		os.Mkdir(folder, 0775)
	}

	_, err1 := os.Stat(folder + "/validator_keys")
	if err1 != nil {
		os.Mkdir(folder, 0775)
	}

	_, errPw := os.Stat(folder + "/passwords")
	if errPw != nil {
		os.Mkdir(folder+"/passwords", 0775)
	}

	mnemonicData := []byte(mnemonic)
	errWrite := os.WriteFile(folder+"/passwords/mnemonic", mnemonicData, 0644)
	if errWrite != nil {
		fmt.Println(folder + "/passwords/mnemonic")
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

	errPW := os.WriteFile(folder+"/passwords/keys", []byte(body.Password), 0644)
	if errPW != nil {
		fmt.Println(errPW)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to create password file")
		return
	}

	command := exec.Command("bash", "-c", shared.BinaryDir+"lukso-deposit-cli/v1.2.6-LUKSO/lukso-deposit-cli "+strings.Join(args, " "))

	if startError := command.Start(); startError != nil {
		log.Fatal(startError)
	}

	command.Wait()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Successfully created keys"); err != nil {
		panic(err)
	}
}

func ImportValidatorKeys(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body importValidatorKeysRequestBody
	errJson := decoder.Decode(&body)
	if errJson != nil {
		panic(errJson)
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
		log.Fatal(startError)
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
	keys, ok := r.URL.Query()["network"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	network := keys[0]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ReadDepositData(network)); err != nil {
		panic(err)
	}
}
