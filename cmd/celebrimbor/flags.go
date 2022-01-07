package main

import (
	"fmt"
	"github.com/lukso-network/lukso-orchestrator/shared/cmd"
	"github.com/urfave/cli/v2"
	"runtime"
)

// Execution layer related flag names
const (
	ELTagFlag         = "el-tag"
	ELDatadirFlag     = "el-datadir"
	ELEthstatsFlag    = "el-ethstats"
	ELBootnodesFlag   = "el-bootnodes"
	ELNetworkIDFlag   = "el-networkid"
	ELPortFlag        = "el-port"
	ELChainIDFlag     = "el-chainid"
	ELHttpApiFlag     = "el-http-apis"
	ELWSApiFlag       = "el-ws-apis"
	ELWSPortFlag      = "el-websocket-port"
	ELEtherbaseFlag   = "el-etherbase"
	ELGenesisFileFlag = "el-genesis"
	ELNotifyFlag      = "el-notify"
	ELVerbosityFlag   = "el-verbosity"
	ELHttpPortFlag    = "el-http-port"
	ELOutputFlag      = "el-output"
	ELWsOriginFlag    = "el-ws-origin"
	ELHttpOriginFlag  = "el-http-origin"
	ELNatFlag         = "EL-nat"

	// Common for CL client(s)
	CLChainConfigFlag = "cl-chain-config"

	// Validator related flag names
	validatorTagFlag                = "validator-tag"
	validatorCLRpcProviderFlag      = "validator-CL-rpc"
	validatorVerbosityFlag          = "validator-verbosity"
	validatorTrustedELFlag          = "validator-trusted-EL"
	validatorWalletPasswordFileFlag = "validator-wallet-password-file"
	validatorDatadirFlag            = "validator-datadir"
	validatorOutputFileFlag         = "validator-output-file"

	// CL related flag names
	CLTagFlag                     = "cl-tag"
	CLGenesisStateFlag            = "cl-genesis-state"
	CLDatadirFlag                 = "cl-datadir"
	CLBootnodesFlag               = "cl-bootnode"
	CLPeerFlag                    = "cl-peer"
	CLOutputFlag                  = "cl-output"
	CLWeb3ProviderFlag            = "cl-web3provider"
	CLDepositContractFlag         = "cl-deposit-contract"
	CLContractDeploymentBlockFlag = "cl-deposit-deployment"
	CLVerbosityFlag               = "cl-verbosity"
	CLMinSyncPeersFlag            = "cl-min-sync-peers"
	CLMaxSyncPeersFlag            = "cl-max-sync-peers"
	CLP2pHostFlag                 = "cl-p2p-host"
	CLP2pLocalFlag                = "cl-p2p-local"
	CLOrcProviderFlag             = "cl-orc-provider"
	CLDisableSyncFlag             = "cl-disable-sync"
	CLOutputFileFlag              = "cl-output-file"
	CLGrpcEndpoint                = "cl-grpc-endpoint"

	DefaultELRPCEndpoint  = "./EL/geth.ipc"
	DefaultCLGRPCEndpoint = "127.0.0.1:4000"
)

var (
	AcceptTOUFlag = &cli.BoolFlag{
		Name:     "accept-terms-of-use",
		Usage:    "you must accept terms of use",
		Required: true,
		Value:    true,
	}
	ForceClearDB = &cli.BoolFlag{
		Name:  "force-clear-db",
		Usage: "Clear any previously stored data at the data directory",
	}
	// LogFileName specifies the log output file name.
	LogFileName = &cli.StringFlag{
		Name:  "log-file",
		Usage: "Specify log file name, relative or absolute",
	}
	appFlags = []cli.Flag{
		AcceptTOUFlag,
		ForceClearDB,
		LogFileName,
	}
	ELFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  ELTagFlag,
			Usage: "provide a tag of EL you would like to run",
			Value: "v0.0.16-alpha",
		},
		&cli.StringFlag{
			Name:  ELDatadirFlag,
			Usage: "provide a path you would like to store your data",
			Value: "./EL",
		},
		&cli.BoolFlag{
			Name:  ELOutputFlag,
			Usage: "do you want to have output attached to your combined output",
			Value: false,
		},
		&cli.StringFlag{
			Name:  ELEthstatsFlag,
			Usage: "nickname:STATS_LOGIN_SECRET@EL_STATS_HOST",
			Value: "",
		},
		&cli.StringFlag{
			Name:  ELBootnodesFlag,
			Usage: "Default value should be ok for test network. Otherwise provide Comma separated enode urls, see at https://geth.ethereum.org/docs/getting-started/private-net.",
			Value: "enode://967db4f56ad0a1a35e3d30632fa600565329a23aff50c9762181810166f3c15b078cca522f930d1a2747778893232336bffd1ea5d2ca60543f1801d4360ea63a@35.204.255.172:0?discport=30301",
		},
		&cli.StringFlag{
			Name:  ELNetworkIDFlag,
			Usage: "provide network id if must be different than default",
			Value: "4004181",
		},
		&cli.StringFlag{
			Name:  ELChainIDFlag,
			Usage: "provide chain id if must be different than default",
			Value: "4004181",
		},
		&cli.StringFlag{
			Name:  ELPortFlag,
			Usage: "provide port for EL",
			Value: "30405",
		},
		&cli.StringFlag{
			Name:  ELHttpApiFlag,
			Usage: "comma separated apis",
			Value: "eth,net",
		},
		&cli.StringFlag{
			Name:  ELHttpPortFlag,
			Usage: "port used in EL http communication",
			Value: "8565",
		},
		&cli.StringFlag{
			Name:  ELWSApiFlag,
			Usage: "comma separated apis",
			Value: "eth,net",
		},
		&cli.StringFlag{
			Name:  ELWSPortFlag,
			Usage: "port for EL api",
			Value: "8546",
		},
		&cli.StringFlag{
			Name:  ELEtherbaseFlag,
			Usage: "your ECDSA public key used to get rewards on EL chain",
			// yes, If you wont set it up, I'll get rewards ;]
			Value: "0x59E3dADc83af3c127a2e29B12B0E86109Bb6d838",
		},
		&cli.StringFlag{
			Name:  ELGenesisFileFlag,
			Usage: "remote genesis file that will be downloaded to spin up the network",
			// yes, If you wont set it up, I'll get rewards ;]
			Value: "https://storage.googleapis.com/l16-common/EL/EL_private_testnet_genesis.json",
		},
		&cli.StringFlag{
			Name:  ELNotifyFlag,
			Usage: "this flag is used to EL engine to notify validator and orchestrator",
			Value: "ws://127.0.0.1:7878,http://127.0.0.1:7877",
		},
		&cli.StringFlag{
			Name:  ELVerbosityFlag,
			Usage: "this flag sets up verobosity for EL",
			Value: "3",
		},
		&cli.StringFlag{
			Name:  ELWsOriginFlag,
			Usage: "this flag sets up websocket accepted origins, default not set",
			Value: "",
		},
		&cli.StringFlag{
			Name:  ELHttpOriginFlag,
			Usage: "this flag sets up http accepted origins, default not set",
			Value: "",
		},
		&cli.StringFlag{
			Name:  ELNatFlag,
			Usage: "this flag sets up http nat to assign static ip for geth, default not set. Example `extip:172.16.254.4`",
			Value: "",
		},
	}
	validatorFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  validatorTagFlag,
			Usage: "provide tag for validator binary. Release must be present in lukso namespace on github",
			Value: "v0.0.18-delta",
		},
		&cli.StringFlag{
			Name:  validatorCLRpcProviderFlag,
			Usage: "provide url without prefix, example: localhost:4000",
			Value: "localhost:4000",
		},
		&cli.StringFlag{
			Name:  CLChainConfigFlag,
			Usage: "path to chain config of CL and validator",
			// TODO: Parse it automatically
			Value: "./CL/v0.0.18-delta/config.yml",
		},
		&cli.BoolFlag{
			Name:  CLOutputFlag,
			Usage: "path to chain config of CL and validator",
			// TODO: Parse it automatically
			Value: false,
		},
		&cli.StringFlag{
			Name:  validatorVerbosityFlag,
			Usage: "provide verbosity of validator",
			Value: "info",
		},
		&cli.StringFlag{
			Name:  validatorTrustedELFlag,
			Usage: "provide host:port for trusted EL, default: http://127.0.0.1:8565",
			Value: "http://127.0.0.1:8565",
		},
		&cli.StringFlag{
			Name:  validatorWalletPasswordFileFlag,
			Usage: "location of file password that you used for generation keys from deposit-cli",
			Value: "./password.txt",
		},
		&cli.StringFlag{
			Name:  validatorDatadirFlag,
			Usage: "location of keys from deposit-cli",
			Value: "",
		},
		&cli.StringFlag{
			Name:  validatorOutputFileFlag,
			Usage: "provide output destination of CL",
			Value: "./validator.log",
		},
	}
	CLFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  CLTagFlag,
			Usage: "provide tag for CL",
			Value: "v0.0.18-delta",
		},
		&cli.StringFlag{
			Name: CLGenesisStateFlag,
			// TODO: see if it is possible to do this via url
			Usage: "provide genesis.ssz file",
			Value: "./CL/v0.0.18-delta/CL_private_testnet_genesis.ssz",
		},
		&cli.StringFlag{
			Name:  CLDatadirFlag,
			Usage: "provide CL datadir",
			Value: "./CL",
		},
		&cli.StringFlag{
			Name:  CLBootnodesFlag,
			Usage: `provide coma separated bootnode enr, default: "enr:-Ku4QANldTRLCRUrY9K4OAmk_ATOAyS_sxdTAaGeSh54AuDJXxOYij1fbgh4KOjD4tb2g3T-oJmMjuJyzonLYW9OmRQBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhAoABweJc2VjcDI1NmsxoQKWfbT1atCho149MGMvpgBWUymiOv9QyXYhgYEBZvPBW4N1ZHCCD6A"`,
			Value: "",
		},
		&cli.StringFlag{
			Name:  CLPeerFlag,
			Usage: `provide coma separated peer enr address, default: ""`,
			Value: "",
		},
		&cli.StringFlag{
			Name:  CLWeb3ProviderFlag,
			Usage: "provide web3 provider (network of deposit contract deployment), default: http://127.0.0.1:8565",
			Value: DefaultELRPCEndpoint,
		},
		&cli.StringFlag{
			Name:  CLDepositContractFlag,
			Usage: "provide deposit contract address",
			Value: "0x000000000000000000000000000000000000cafe",
		},
		&cli.StringFlag{
			Name:  CLContractDeploymentBlockFlag,
			Usage: "provide deployment height of deposit contract, default 0.",
			Value: "0",
		},
		&cli.StringFlag{
			Name:  CLVerbosityFlag,
			Usage: "provide verobosity for CL",
			Value: "info",
		},
		&cli.StringFlag{
			Name:  CLMinSyncPeersFlag,
			Usage: "provide min sync peers for CL, default 0",
			Value: "0",
		},
		&cli.StringFlag{
			Name:  CLMaxSyncPeersFlag,
			Usage: "provide max sync peers for CL, default 25",
			Value: "25",
		},
		&cli.StringFlag{
			Name:  CLP2pHostFlag,
			Usage: "provide p2p host for CL, default empty",
			Value: "",
		}, &cli.StringFlag{
			Name:  CLP2pLocalFlag,
			Usage: "provide p2p local ip for CL, default empty",
			Value: "",
		},
		&cli.StringFlag{
			Name:  CLOrcProviderFlag,
			Usage: "provide orchestrator provider, default http://127.0.0.1:7878",
			Value: "http://127.0.0.1:7877",
		},
		&cli.BoolFlag{
			Name:  CLDisableSyncFlag,
			Usage: "disable initial sync phase",
			Value: false,
		},
		&cli.StringFlag{
			Name:  CLOutputFileFlag,
			Usage: "provide output destination of CL",
			Value: "./CL.log",
		},
	}
)

// setupOperatingSystem will parse flags and use it to deduce which system dependencies are required
func setupOperatingSystem() {
	systemOs = runtime.GOOS
}

func prepareCLFlags(ctx *cli.Context) (CLArguments []string) {
	if !ctx.Bool(AcceptTOUFlag.Name) {
		log.Fatal("you must accept terms of use")
		ctx.Done()

		return
	}

	CLArguments = append(CLArguments, "--accept-terms-of-use")

	if ctx.IsSet(cmd.ForceClearDB.Name) {
		CLArguments = append(CLArguments, "--force-clear-db")
	}

	CLArguments = append(CLArguments, fmt.Sprintf("--chain-id=%s", ctx.String(ELChainIDFlag)))
	CLArguments = append(
		CLArguments,
		fmt.Sprintf("--network-id=%s", ctx.String(ELNetworkIDFlag)))
	CLArguments = append(CLArguments, fmt.Sprintf("--datadir"))
	CLArguments = append(CLArguments, ctx.String(CLDatadirFlag))

	// This flag can be shared for sure. There is no possibility to use different specs for validator and CL.
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--chain-config-file=%s",
		ctx.String(CLChainConfigFlag),
	))
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--bootstrap-node=%s",
		ctx.String(CLBootnodesFlag),
	))

	if "" != ctx.String(CLPeerFlag) {
		CLArguments = append(CLArguments, fmt.Sprintf(
			"--peer=%s",
			ctx.String(CLPeerFlag),
		))
	}

	if ctx.Bool(CLDisableSyncFlag) {
		CLArguments = append(CLArguments, "--disable-sync")
	}

	CLArguments = append(CLArguments, fmt.Sprintf(
		"--http-web3provider=%s",
		ctx.String(CLWeb3ProviderFlag),
	))
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--deposit-contract=%s",
		ctx.String(CLDepositContractFlag),
	))
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--contract-deployment-block=%s",
		ctx.String(CLContractDeploymentBlockFlag),
	))
	CLArguments = append(CLArguments, "--rpc-host=0.0.0.0")
	CLArguments = append(CLArguments, "--monitoring-host=0.0.0.0")
	CLArguments = append(CLArguments, "--verbosity")
	CLArguments = append(CLArguments, ctx.String(CLVerbosityFlag))
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--min-sync-peers=%s",
		ctx.String(CLMinSyncPeersFlag),
	))
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--p2p-max-peers=%s",
		ctx.String(CLMaxSyncPeersFlag),
	))

	if "" != ctx.String(CLP2pHostFlag) {
		CLArguments = append(CLArguments, fmt.Sprintf(
			"--p2p-host-ip=%s",
			ctx.String(CLP2pHostFlag),
		))
	}

	if "" != ctx.String(CLP2pLocalFlag) {
		CLArguments = append(CLArguments, fmt.Sprintf(
			"--p2p-local-ip=%s",
			ctx.String(CLP2pLocalFlag),
		))
	}

	CLArguments = append(CLArguments, fmt.Sprintf(
		"--log-file=%s",
		ctx.String(CLOutputFileFlag),
	))
	CLArguments = append(CLArguments, "--lukso-network")
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--orc-http-provider=%s",
		ctx.String(CLOrcProviderFlag),
	))
	CLArguments = append(CLArguments, fmt.Sprintf(
		"--genesis-state=%s",
		ctx.String(CLGenesisStateFlag),
	))

	return
}

func prepareValidatorFlags(ctx *cli.Context) (validatorArguments []string) {
	if !ctx.Bool(AcceptTOUFlag.Name) {
		log.Fatal("you must accept terms of use")
		ctx.Done()

		return
	}

	validatorArguments = append(validatorArguments, "--accept-terms-of-use")

	if ctx.IsSet(ForceClearDB.Name) {
		validatorArguments = append(validatorArguments, "--force-clear-db")
	}

	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--chain-config-file=%s",
		ctx.String(CLChainConfigFlag),
	))
	validatorArguments = append(validatorArguments, "--verbosity")
	validatorArguments = append(validatorArguments, ctx.String(validatorVerbosityFlag))
	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--EL-http-provider=%s",
		ctx.String(validatorTrustedELFlag),
	))
	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--log-file=%s",
		ctx.String(validatorOutputFileFlag),
	))
	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--wallet-password-file=%s",
		ctx.String(validatorWalletPasswordFileFlag),
	))
	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--datadir=%s",
		ctx.String(CLDatadirFlag),
	))

	if "" != ctx.String(validatorDatadirFlag) {
		validatorArguments = append(validatorArguments, fmt.Sprintf(
			"--wallet-dir=%s",
			ctx.String(validatorDatadirFlag),
		))
	}

	validatorArguments = append(validatorArguments, "--lukso-network")

	return
}

func prepareELFlags(ctx *cli.Context) (ELArguments []string) {
	ELArguments = append(ELArguments, "--datadir")
	ELArguments = append(ELArguments, ctx.String(ELDatadirFlag))

	if len(ctx.String(ELEthstatsFlag)) > 1 {
		ELArguments = append(ELArguments, "--ethstats")
		ELArguments = append(ELArguments, ctx.String(ELEthstatsFlag))
	}

	if len(ctx.String(ELBootnodesFlag)) > 1 {
		ELArguments = append(ELArguments, "--bootnodes")
		ELArguments = append(ELArguments, ctx.String(ELBootnodesFlag))
	}

	ELArguments = append(ELArguments, "--networkid")
	ELArguments = append(ELArguments, ctx.String(ELNetworkIDFlag))
	ELArguments = append(ELArguments, "--port")
	ELArguments = append(ELArguments, ctx.String(ELPortFlag))

	// Http api
	// TODO: change to new --http, because -rpc is deprecated in EL
	ELArguments = append(ELArguments, "--rpc")
	ELArguments = append(ELArguments, "--rpcaddr")
	ELArguments = append(ELArguments, "0.0.0.0")
	ELArguments = append(ELArguments, "--rpcapi")
	ELArguments = append(ELArguments, ctx.String(ELHttpApiFlag))
	ELArguments = append(ELArguments, "--rpcport")
	ELArguments = append(ELArguments, ctx.String(ELHttpPortFlag))

	if "" != ctx.String(ELHttpOriginFlag) {
		ELArguments = append(ELArguments, "--http.corsdomain")
		ELArguments = append(ELArguments, ctx.String(ELHttpOriginFlag))
	}

	// Nat extIP
	if "" != ctx.String(ELNatFlag) {
		ELArguments = append(ELArguments, "--nat")
		ELArguments = append(ELArguments, ctx.String(ELNatFlag))
	}

	// Websocket
	ELArguments = append(ELArguments, "--ws")
	ELArguments = append(ELArguments, "--ws.addr")
	ELArguments = append(ELArguments, "0.0.0.0")
	ELArguments = append(ELArguments, "--ws.api")
	ELArguments = append(ELArguments, ctx.String(ELWSApiFlag))
	ELArguments = append(ELArguments, "--ws.port")
	ELArguments = append(ELArguments, ctx.String(ELWSPortFlag))

	if "" != ctx.String(ELWsOriginFlag) {
		ELArguments = append(ELArguments, "--ws.origins")
		ELArguments = append(ELArguments, ctx.String(ELWsOriginFlag))
	}

	// Miner
	ELArguments = append(ELArguments, "--miner.etherbase")
	ELArguments = append(ELArguments, ctx.String(ELEtherbaseFlag))
	ELArguments = append(ELArguments, "--mine")
	ELArguments = append(ELArguments, "--miner.notify")
	ELArguments = append(ELArguments, ctx.String(ELNotifyFlag))

	// Verbosity
	ELArguments = append(ELArguments, "--verbosity")
	ELArguments = append(ELArguments, ctx.String(ELVerbosityFlag))

	return
}
