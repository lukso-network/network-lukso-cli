# LUKSO CLI
>⚠️ This page may change. Not everything is ready yet.


## Repository struct
* `./cmd/lukso-cli` is currently used to store GO source files for this `cli`
* `./test` contains all test configs used to test this script with the results stored in `./test/logs` and example Linux binary under `./test/bin` directory
## Glossary

`validator node` - it's a setup of `geth`, `beacon-chain` and `validator` binaries

`lukso-cli` is default name of `cli` binary

## Running

### Archive node
1. Build `go build ./cmd/lukso-cli`

2. Run `./lukso-cli start arch` to start a validator node

### Validator node
1. Build `go build ./cmd/lukso-cli`

2. Create `CL-Validator-wallet` directory in project root

3. Insert all pre-generated `prysm` wallet files into `CL-Validator-wallet` directory:
```
./CL-Validator-wallet
├── direct
│   └── accounts
│       └── all-accounts.keystore.json
└── keymanageropts.json
```

4. Create `password.txt` text file in project root and fill with proper `prysm` wallet password. Default password is: `Test1234`

5. Run `./lukso-cli start all` to start a validator node

## Stopping

Run `./lukso-cli stop` to stop a validator node
