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

type getDepositDataRequestBody struct {
	Network string
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

	// errDeleteValidatoKeys := os.Remove(folder + "/validator_keys")
	// if errDeleteValidatoKeys != nil {
	// 	log.Fatal(errDeleteValidatoKeys)
	// }

	// errDeleteVanguardWallet := os.Remove(folder + "/vanguard_wallet")
	// if errDeleteVanguardWallet != nil {
	// 	log.Fatal(errDeleteVanguardWallet)
	// }

	err := os.Chmod(folder, 0775)
	if err != nil {
		os.Mkdir(folder, 0775)
	}

	fmt.Println(folder + "/passwords/mnemonic")

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

	fmt.Println(args)
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
		log.Printf(in.Text()) // write each line to your log, or anything you need
	}

	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}

	command.Wait()

	zipFile := zipKeys(body.Network)

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("validator_keys.zip"))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, zipFile)
}

func GetDepositData(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["network"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	network := keys[0]

	fmt.Println(ReadDepositData(network))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ReadDepositData(network)); err != nil {
		panic(err)
	}
}
