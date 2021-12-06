package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type LuksoValues struct {
	Force              bool   `yaml:"FORCE"`
	Network            string `yaml:"NETWORK"`
	Validate           bool   `yaml:"VALIDATE"`
	Coinbase           string `yaml:"COINBASE"`
	NodeName           string `yaml:"NODE_NAME"`
	DataDir            string `yaml:"DATADIR"`
	LogsDir            string `yaml:"LOGSDIR"`
	KeysDir            string `yaml:"KEYSDIR"`
	KeysPasswordFile   string `yaml:"KEYS_PASSWORD_FILE"`
	WalletDir          string `yaml:"WALLET_DIR"`
	WalletPasswordFile string `yaml:"WALLET_PASSWORD_FILE"`

	Orchestrator struct {
		Tag                 string `yaml:"TAG"`
		Verbosity           string `yaml:"VERBOSITY"`
		VanguardRPCEndpoint string `yaml:"VanguardRPCEndpoint"`
		PandoraRPCEndpoint  string `yaml:"PandoraRPCEndpoint"`
	} `yaml:"ORCHESTRATOR"`

	Pandora struct {
		Tag                 string `yaml:"TAG"`
		Verbosity           string `yaml:"VERBOSITY"`
		Bootnodes           string `yaml:"BOOTNODES"`
		Port                int    `yaml:"PORT"`
		HttpAddr            string `yaml:"HTTP_ADDR"`
		HttpPort            int    `yaml:"HTTP_PORT"`
		WebsocketsAddr      string `yaml:"WS_ADDR"`
		WebsocketsPort      int    `yaml:"WS_PORT"`
		HTTPMinerAddr       string `yaml:"HTTP_MINER_ADDR"`
		WebsocketsMinerAddr string `yaml:"WS_MINER_ADDR"`
		Ethstats            string `yaml:"ETHSTATS"`
		// TODO: find different name
		UPExpose     bool `yaml:"UP_EXPOSE"`
		UnsafeExpose bool `yaml:"UNSAFE_EXPOSE"`
	} `yaml:"PANDORA"`

	Vanguard struct {
		Tag                     string `yaml:"TAG"`
		Verbosity               string `yaml:"VERBOSITY"`
		Bootnodes               string `yaml:"BOOTNODES"`
		P2PPrivKEY              string `yaml:"P2P_PRIVKEY"`
		ExternalIP              string `yaml:"EXTERNAL_IP"`
		P2PHostDNS              string `yaml:"P2PHostDNS"`
		RPCHost                 string `yaml:"RPCHost"`
		RPCPort                 int    `yaml:"RPCPort"`
		UDPPort                 int    `yaml:"UDPPort"`
		TCPPort                 int    `yaml:"TCPPort"`
		MonitoringHost          string `yaml:"MONITORING_HOST"`
		HTTPWeb3Provider        string `yaml:"HTTP_WEB3_PROVIDER"`
		GRPCGatewayPort         int    `yaml:"GRPC_GATEWAY_PORT"`
		OrchestratorRPCProvider string `yaml:"ORCHESTRATOR_RPC_PROVIDER"`
		MinSyncPeers            int    `yaml:"MIN_SYNC_PEERS"`
		MaxP2PPeers             int    `yaml:"MAX_P2P_PEERS"`
		Ethstats                string `yaml:"ETHSTATS"`
		EthstatsMetrics         string `yaml:"ETHSTATS_METRICS"`
	}

	Validator struct {
		Tag                 string `yaml:"TAG"`
		Verbosity           string `yaml:"VERBOSITY"`
		BeaconRPCProvider   string `yaml:"BEACON_RPC_PROVIDER"`
		PandoraHTTPProvider string `yaml:"PANDORA_HTTP_PROVIDER"`
	}
}

type NetworkValues struct {
	ChainID           int    `yaml:"CHAIN_ID"`
	NetworkID         int    `yaml:"NETWORK_ID"`
	ForkChoice        int    `yaml:"FORK_CHOICE"`
	PandoraBootnodes  string `yaml:"PANDORA_BOOTNODES"`
	VanguardBootnodes string `yaml:"VANGUARD_BOOTNODES"`
}

func LoadNetworkConfig(Network string) (error, NetworkValues) {
	var NetworkConfig NetworkValues
	configFilePath := "/opt/lukso/networks/" + Network + "/config/network-config.yaml"

	buf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err, NetworkConfig
	}

	c := &NetworkConfig
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("in file %q: %v", configFilePath, err), NetworkConfig
	}

	return nil, NetworkConfig
}

func LoadDefaults(LuksoSettings *LuksoValues) {
	LuksoSettings.Force = false
	LuksoSettings.Network = "l15"
	LuksoSettings.Validate = false
	LuksoSettings.Coinbase = ""
	LuksoSettings.NodeName = ""

	homeDir, err := os.UserHomeDir()
	if err != nil {
		LuksoSettings.DataDir = homeDir + "/.lukso/" + LuksoSettings.Network + "/datadir"
		LuksoSettings.LogsDir = homeDir + "/.lukso/" + LuksoSettings.Network + "/logs"
		LuksoSettings.KeysDir = homeDir + "/.lukso/" + LuksoSettings.Network + "/validator_keys"
		LuksoSettings.WalletDir = homeDir + "/.lukso/" + LuksoSettings.Network + "/wallet"
	}

	LuksoSettings.KeysPasswordFile = ""
	LuksoSettings.WalletPasswordFile = ""

	LuksoSettings.Orchestrator.Tag = ""
	LuksoSettings.Orchestrator.Verbosity = "info"
	LuksoSettings.Orchestrator.VanguardRPCEndpoint = "127.0.0.1:4000"
	LuksoSettings.Orchestrator.PandoraRPCEndpoint = "ws://127.0.0.1:8546"

	LuksoSettings.Pandora.Tag = ""
	LuksoSettings.Pandora.Verbosity = "info"
	LuksoSettings.Pandora.Bootnodes = ""
	LuksoSettings.Pandora.Port = 30303
	LuksoSettings.Pandora.HttpAddr = "127.0.0.1"
	LuksoSettings.Pandora.HttpPort = 8545
	LuksoSettings.Pandora.WebsocketsAddr = "127.0.0.1"
	LuksoSettings.Pandora.WebsocketsPort = 8546
	LuksoSettings.Pandora.Ethstats = ""
	LuksoSettings.Pandora.UPExpose = false
	LuksoSettings.Pandora.UnsafeExpose = false

	LuksoSettings.Vanguard.Tag = ""
	LuksoSettings.Vanguard.Verbosity = "info"
	LuksoSettings.Vanguard.Bootnodes = ""
	LuksoSettings.Vanguard.P2PPrivKEY = ""
	LuksoSettings.Vanguard.ExternalIP = ""
	LuksoSettings.Vanguard.P2PHostDNS = ""
	LuksoSettings.Vanguard.RPCHost = "127.0.0.1"
	LuksoSettings.Vanguard.RPCPort = 4000
	LuksoSettings.Vanguard.UDPPort = 12000
	LuksoSettings.Vanguard.TCPPort = 13000
	LuksoSettings.Vanguard.MonitoringHost = "127.0.0.1"
	LuksoSettings.Vanguard.HTTPWeb3Provider = "http://127.0.0.1:8545"
	LuksoSettings.Vanguard.GRPCGatewayPort = 3500
	LuksoSettings.Vanguard.OrchestratorRPCProvider = "http://127.0.0.1:7877"
	LuksoSettings.Vanguard.MinSyncPeers = 2
	LuksoSettings.Vanguard.MaxP2PPeers = 50
	LuksoSettings.Vanguard.Ethstats = ""
	LuksoSettings.Vanguard.EthstatsMetrics = ""

	LuksoSettings.Validator.Tag = ""
	LuksoSettings.Validator.Verbosity = "info"
	LuksoSettings.Validator.BeaconRPCProvider = "127.0.0.1:4000"
	LuksoSettings.Validator.PandoraHTTPProvider = "http://127.0.0.1:8545"

}
