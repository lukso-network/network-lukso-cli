package config

type LuksoValues struct {
	Force              bool   `yaml:"FORCE"`
	Network            string `yaml:"NETWORK"`
	Validate           bool   `yaml:"VALIDATE"`
	Coinbase           string `yaml:"COINBASE"`
	NodeName           string `yaml:"NODE_NAME"`
	LogsDir            string `yaml:"LOGSDIR"`
	DataDir            string `yaml:"DATADIR"`
	KeysDir            string `yaml:"KEYSDIR"`
	KeysPassFile       string `yaml:"KEYS_PASSWORD_FILE"`
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
		HttpPort            int    `yaml:"HTTP_PORT"`
		HttpAddr            string `yaml:"HTTP_ADDR"`
		WebsocketsPort      int    `yaml:"WS_PORT"`
		WebsocketsAddr      string `yaml:"WS_ADDR"`
		BeaconRPCProvider   string `yaml:"BEACON_RPC_PROVIDER"`
		PandoraHTTPProvider string `yaml:"PANDORA_HTTP_PROVIDER"`
	}
}

func LoadDefaults(LuksoSettings *LuksoValues) {
	LuksoSettings.Force = false
	LuksoSettings.Network = "l15"
	LuksoSettings.Orchestrator.Verbosity = "info"
	LuksoSettings.Pandora.Verbosity = "info"
}
