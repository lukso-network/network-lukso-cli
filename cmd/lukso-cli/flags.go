package main

import (
	"fmt"
	"github.com/lukso-network/lukso-orchestrator/shared/cmd"
	"github.com/urfave/cli/v2"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Execution layer related flag names
const (
	nodeNameFlag = "node-name"

	ELTagFlag        = "el-tag"
	ELDatadirFlag    = "el-datadir"
	ELEthstatsFlag   = "el-ethstats"
	ELBootnodesFlag  = "el-bootnodes"
	ELNetworkIDFlag  = "el-networkid"
	ELPortFlag       = "el-port"
	ELChainIDFlag    = "el-chainid"
	ELHttpApiFlag    = "el-http-apis"
	ELWSApiFlag      = "el-ws-apis"
	ELWSPortFlag     = "el-websocket-port"
	ELEtherbaseFlag  = "el-etherbase"
	ELVerbosityFlag  = "el-verbosity"
	ELHttpPortFlag   = "el-http-port"
	ELOutputFlag     = "el-output"
	ELWsOriginFlag   = "el-ws-origin"
	ELHttpOriginFlag = "el-http-origin"
	ELNatFlag        = "el-nat"
	ELLogFileFlag    = "el-log-file"

	// CLChainConfigFlag Common for CL client(s)
	CLChainConfigFlag = "cl-chain-config"

	// Validator related flag names
	validatorTagFlag                = "validator-tag"
	validatorCLRpcProviderFlag      = "validator-CL-rpc"
	validatorVerbosityFlag          = "validator-verbosity"
	validatorWalletPasswordFileFlag = "validator-wallet-password-file"
	validatorDatadirFlag            = "validator-datadir"
	validatorWalletDatadirFlag      = "validator-wallet-datadir"
	validatorLogFileFlag            = "validator-log-file"
	validatorOutputFlag             = "validator-output"

	// CLTagFlag CL related flag names
	CLTagFlag                     = "cl-tag"
	CLGenesisStateFlag            = "cl-genesis-state"
	CLDatadirFlag                 = "cl-datadir"
	CLBootnodesFlag               = "cl-bootnodes"
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
	CLP2PTCPPort                  = "cl-p2p-port-tcp"
	CLP2PUDPPort                  = "cl-p2p-port-udp"
	CLGRPCGatewayPort             = "cl-grpc-gateway-port"
	CLRPCPort                     = "cl-rpc-port"
	CLDisableSyncFlag             = "cl-disable-sync"
	CLLogFileFlag                 = "cl-log-file"

	// CL Stats Client related flag names
	CLStatsClientFlag        = "cl-stats-tag"
	clStatsClientDatadirFlag = "cl-stats-datadir"
	clStatsOutputFlag        = "cl-stats-output"

	DefaultELRPCEndpoint        = "http://127.0.0.1:8598"
	DefaultLogFilenameSeparator = "_"
)

var (
	DefaultNodeName = fmt.Sprintf("local-node-%d", time.Now().Unix())
	CLGrpcEndpoint  = fmt.Sprintf("127.0.0.1:%d", DefaultCLGRPCPort)
	ForceClearDB    = &cli.BoolFlag{
		Name:  "force-clear-db",
		Usage: "Clear any previously stored data at the data directory",
	}
	// LogFileName specifies the log output file name.
	LogFileName = &cli.StringFlag{
		Name:  "log-file",
		Usage: "Specify log file name, relative or absolute",
	}
	NodeNameFlag = &cli.StringFlag{
		Name:  nodeNameFlag,
		Usage: "Specify your node name, this is how you recognize your node on the network stats pages",
		Value: DefaultNodeName,
	}
	appFlags = []cli.Flag{
		ForceClearDB,
		LogFileName,
		NodeNameFlag,
	}
	ELFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  ELTagFlag,
			Usage: "provide a tag of EL you would like to run",
			Value: "v1.0.0",
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
			Value: "",
		},
		&cli.StringFlag{
			Name:  ELNetworkIDFlag,
			Usage: "provide network id if must be different than default",
			Value: strconv.Itoa(DefaultELNetworkID),
		},
		&cli.StringFlag{
			Name:  ELChainIDFlag,
			Usage: "provide chain id if must be different than default",
			Value: strconv.Itoa(DefaultELNetworkID),
		},
		&cli.StringFlag{
			Name:  ELPortFlag,
			Usage: "provide port for EL",
			Value: strconv.Itoa(DefaultELP2PPort),
		},
		&cli.StringFlag{
			Name:  ELHttpApiFlag,
			Usage: "comma separated apis",
			Value: "engine,net,eth,admin,debug",
		},
		&cli.StringFlag{
			Name:  ELHttpPortFlag,
			Usage: "port used in EL http communication",
			Value: strconv.Itoa(DefaultELHTTPPort),
		},
		&cli.StringFlag{
			Name:  ELWSApiFlag,
			Usage: "comma separated apis",
			Value: "engine,net,eth,admin,debug",
		},
		&cli.StringFlag{
			Name:  ELWSPortFlag,
			Usage: "port for EL api",
			Value: strconv.Itoa(DefaultELWSPort),
		},
		&cli.StringFlag{
			Name:  ELEtherbaseFlag,
			Usage: "your ECDSA public key used to get rewards on EL chain",
			// yes, If you wont set it up, I'll get rewards ;]
			Value: "0x59E3dADc83af3c127a2e29B12B0E86109Bb6d838",
		},
		&cli.StringFlag{
			Name:  ELVerbosityFlag,
			Usage: "this flag sets up verobosity for EL",
			Value: "3",
		},
		&cli.StringFlag{
			Name:  ELWsOriginFlag,
			Usage: "this flag sets up websocket accepted origins, default localhost",
			Value: "localhost",
		},
		&cli.StringFlag{
			Name:  ELHttpOriginFlag,
			Usage: "this flag sets up http accepted origins, default localhost",
			Value: "localhost",
		},
		&cli.StringFlag{
			Name:  ELNatFlag,
			Usage: "this flag sets up http nat to assign static ip for geth, default not set. Example `extip:172.16.254.4`",
			Value: "",
		},
		&cli.StringFlag{
			Name:  ELLogFileFlag,
			Usage: "provide output destination of EL",
			Value: "./EL.log",
		},
	}
	validatorFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  validatorTagFlag,
			Usage: "provide tag for validator binary. Release must be present in lukso namespace on github",
			Value: "v1.0.0",
		},
		&cli.StringFlag{
			Name:  validatorCLRpcProviderFlag,
			Usage: fmt.Sprintf("provide url without prefix, example: 127.0.0.1:%d", DefaultCLGRPCPort),
			Value: fmt.Sprintf("127.0.0.1:%d", DefaultCLGRPCPort),
		},
		&cli.StringFlag{
			Name:  CLChainConfigFlag,
			Usage: "path to chain config of CL and validator",
			// TODO: Parse it automatically
			Value: fmt.Sprintf("./CL/v1.0.0/%s", CLConfigDependencyName),
		},
		&cli.BoolFlag{
			Name:  validatorOutputFlag,
			Usage: "do you want to have output attached to your combined output",
			Value: false,
		},
		&cli.StringFlag{
			Name:  validatorVerbosityFlag,
			Usage: "provide verbosity of validator",
			Value: "debug",
		},
		&cli.StringFlag{
			Name:  validatorWalletPasswordFileFlag,
			Usage: "location of file password that you used for generation keys from deposit-cli",
			Value: "./password.txt",
		},
		&cli.StringFlag{
			Name:  validatorDatadirFlag,
			Usage: "location of CL Validator database files",
			Value: "./CL-Validator",
		},
		&cli.StringFlag{
			Name:  validatorWalletDatadirFlag,
			Usage: "location of keys from deposit-cli - validator wallet dir",
			Value: "./CL-Validator-wallet",
		},
		&cli.StringFlag{
			Name:  validatorLogFileFlag,
			Usage: "provide output destination of CL",
			Value: "./CL-Validator.log",
		},
	}
	CLFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  CLTagFlag,
			Usage: "provide tag for CL",
			Value: "v1.0.0",
		},
		&cli.StringFlag{
			Name: CLGenesisStateFlag,
			// TODO: see if it is possible to do this via url
			Usage: "provide genesis.ssz file",
			Value: fmt.Sprintf("./CL/v1.0.0/%s", CLGenesisDependencyName),
		},
		&cli.StringFlag{
			Name:  CLDatadirFlag,
			Usage: "provide CL datadir",
			Value: "./CL",
		},
		&cli.StringFlag{
			Name:  CLBootnodesFlag,
			Usage: `provide coma separated bootnode enr, default: "enr:-Ku4QANldTRLCRUrY9..."`,
			Value: "",
		},
		&cli.StringFlag{
			Name:  CLPeerFlag,
			Usage: `provide coma separated peer enr address, default: ""`,
			Value: "",
		},
		&cli.StringFlag{
			Name:  CLWeb3ProviderFlag,
			Usage: "provide web3 provider (network of deposit contract deployment), default: IPC",
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
			Value: "debug",
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
		},
		&cli.StringFlag{
			Name:  CLP2pLocalFlag,
			Usage: "provide p2p local ip for CL, default empty",
			Value: "",
		},
		&cli.StringFlag{
			Name:  CLP2PTCPPort,
			Usage: fmt.Sprintf("provide p2p port for tcp, default: %d", DefaultCLP2PTCPPort),
			Value: fmt.Sprintf("%d", DefaultCLP2PTCPPort),
		},
		&cli.StringFlag{
			Name:  CLP2PUDPPort,
			Usage: fmt.Sprintf("provide p2p port for udp, default: %d", DefaultCLP2PUDPPort),
			Value: fmt.Sprintf("%d", DefaultCLP2PUDPPort),
		},
		&cli.StringFlag{
			Name:  CLGRPCGatewayPort,
			Usage: fmt.Sprintf("provide p2p port for udp, default: %d", DefaultCLGRPCGatewayPort),
			Value: fmt.Sprintf("%d", DefaultCLGRPCGatewayPort),
		},
		&cli.StringFlag{
			Name:  CLRPCPort,
			Usage: fmt.Sprintf("provide p2p port for udp, default: %d", DefaultCLGRPCPort),
			Value: fmt.Sprintf("%d", DefaultCLGRPCPort),
		},
		&cli.BoolFlag{
			Name:  CLDisableSyncFlag,
			Usage: "disable initial sync phase",
			Value: false,
		},
		&cli.StringFlag{
			Name:  CLLogFileFlag,
			Usage: "provide output destination of CL",
			Value: "./CL.log",
		},
		&cli.BoolFlag{
			Name:  CLOutputFlag,
			Usage: "do you want to have output attached to your combined output",
			Value: false,
		},
	}
	CLStatsClientFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  CLStatsClientFlag,
			Usage: "provide tag for CL Stats Client",
			Value: "v1.0.0",
		},
		&cli.StringFlag{
			Name:  clStatsClientDatadirFlag,
			Usage: "location of CL Stats Client",
			Value: "./CL-Stats-Client",
		},
		&cli.BoolFlag{
			Name:  clStatsOutputFlag,
			Usage: "do you want to have output attached to your combined output",
			Value: false,
		},
	}
)

// setupOperatingSystem will parse flags and use it to deduce which system dependencies are required
func setupOperatingSystem() {
	systemOs = runtime.GOOS
	systemArch = runtime.GOARCH
}

func prepareCLFlags(ctx *cli.Context) (CLArguments []string) {
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

	if len(ctx.String(CLBootnodesFlag)) > 1 {
		bootstrapNodes := strings.Split(ctx.String(CLBootnodesFlag), ",")
		for _, enr := range bootstrapNodes {
			CLArguments = append(CLArguments, fmt.Sprintf(
				"--bootstrap-node=%s",
				enr,
			))
		}
	}

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
		fmt.Sprintf("%s%s%d", ctx.String(CLLogFileFlag), DefaultLogFilenameSeparator, time.Now().Unix()),
	))

	CLArguments = append(CLArguments, fmt.Sprintf(
		"--genesis-state=%s",
		ctx.String(CLGenesisStateFlag),
	))

	CLArguments = append(CLArguments, fmt.Sprintf(
		"--p2p-tcp-port=%s",
		ctx.String(CLP2PTCPPort),
	))

	CLArguments = append(CLArguments, fmt.Sprintf(
		"--p2p-udp-port=%s",
		ctx.String(CLP2PUDPPort),
	))

	CLArguments = append(CLArguments, "--enable-debug-rpc-endpoints")

	CLArguments = append(CLArguments, "--kintsugi-testnet")

	CLArguments = append(CLArguments, "--grpc-gateway-port")
	CLArguments = append(CLArguments, ctx.String(CLGRPCGatewayPort))

	CLArguments = append(CLArguments, fmt.Sprintf(
		"--rpc-port=%s",
		ctx.String(CLRPCPort),
	))

	fmt.Println("CLArguments")
	fmt.Println(CLArguments)

	return
}

func prepareValidatorFlags(ctx *cli.Context) (validatorArguments []string) {
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
		"--log-file=%s",
		fmt.Sprintf("%s%s%d", ctx.String(validatorLogFileFlag), DefaultLogFilenameSeparator, time.Now().Unix()),
	))
	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--wallet-password-file=%s",
		ctx.String(validatorWalletPasswordFileFlag),
	))
	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--datadir=%s",
		ctx.String(validatorDatadirFlag),
	))

	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--wallet-dir=%s",
		ctx.String(validatorWalletDatadirFlag),
	))

	validatorArguments = append(validatorArguments, fmt.Sprintf(
		"--beacon-rpc-provider=%s",
		ctx.String(validatorCLRpcProviderFlag),
	))

	return
}

func prepareCLStatsClientFlags(ctx *cli.Context) (clStatsClientArguments []string) {
	clStatsClientArguments = append(clStatsClientArguments, "run")
	clStatsClientArguments = append(clStatsClientArguments, "--beacon.type=prysm")
	clStatsClientArguments = append(clStatsClientArguments, fmt.Sprintf("--eth2stats.node-name=%s", ctx.String(nodeNameFlag)))
	clStatsClientArguments = append(clStatsClientArguments, "--eth2stats.addr=34.147.116.58:9090")
	clStatsClientArguments = append(clStatsClientArguments, "--eth2stats.tls=false")
	clStatsClientArguments = append(clStatsClientArguments, "--beacon.metrics-addr=http://127.0.0.1:8080/metrics")
	clStatsClientArguments = append(clStatsClientArguments, fmt.Sprintf("--beacon.addr=127.0.0.1:%d", DefaultCLGRPCPort))

	return
}

func prepareELFlags(ctx *cli.Context) (ELArguments []string) {
	ELArguments = append(ELArguments, "--datadir")
	ELArguments = append(ELArguments, ctx.String(ELDatadirFlag))
	ELArguments = append(ELArguments, "--datadir.ancient")
	ELArguments = append(ELArguments, ctx.String(ELDatadirFlag))
	ELArguments = append(ELArguments, "--ethash.cachedir")
	ELArguments = append(ELArguments, ctx.String(ELDatadirFlag))

	ethstatsArguments := []string{
		"--ethstats",
		fmt.Sprintf("%s:@dev.stats.pandora.l15.lukso.network", ctx.String(nodeNameFlag)),
	}

	ELArguments = append(ELArguments, ethstatsArguments...)

	if len(ctx.String(ELBootnodesFlag)) > 1 {
		ELArguments = append(ELArguments, "--bootnodes")
		ELArguments = append(ELArguments, ctx.String(ELBootnodesFlag))
	}

	ELArguments = append(ELArguments, "--networkid")
	ELArguments = append(ELArguments, ctx.String(ELNetworkIDFlag))
	ELArguments = append(ELArguments, "--port")
	ELArguments = append(ELArguments, ctx.String(ELPortFlag))

	// Http api
	ELArguments = append(ELArguments, "--http")
	ELArguments = append(ELArguments, "--http.addr")
	ELArguments = append(ELArguments, "0.0.0.0")
	ELArguments = append(ELArguments, "--http.api")
	ELArguments = append(ELArguments, ctx.String(ELHttpApiFlag))
	ELArguments = append(ELArguments, "--http.port")
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

	// Catalyst
	ELArguments = append(ELArguments, "--catalyst")

	// Verbosity
	ELArguments = append(ELArguments, "--verbosity")
	ELArguments = append(ELArguments, ctx.String(ELVerbosityFlag))

	fmt.Println("ELArguments")
	fmt.Println(ELArguments)

	return
}
