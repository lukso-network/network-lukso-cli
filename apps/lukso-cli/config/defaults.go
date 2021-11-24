package config

type LuksoValues struct {
	Force    bool   `yaml:"FORCE"`
	Network  string `yaml:"NETWORK"`
	Coinbase string `yaml:"COINBASE"`
	NodeName string `yaml:"NODE_NAME"`
	LogsDir  string `yaml:"LOGSDIR"`
	DataDir  string `yaml:"DATADIR"`

	Orchestrator struct {
		Tag       string `yaml:"TAG"`
		Verbosity string `yaml:"VERBOSITY"`
	} `yaml:"ORCHESTRATOR"`

	Pandora struct {
		Tag       string `yaml:"TAG"`
		Verbosity string `yaml:"VERBOSITY"`
		Bootnodes string `yaml:"BOOTNODES"`
		HttpPort  string `yaml:"HTTP_PORT"`
		// TODO: find different name
		UPExpose     bool `yaml:"UP_EXPOSE"`
		UnsafeExpose bool `yaml:"UNSAFE_EXPOSE"`
	} `yaml:"PANDORA"`

	Vanguard struct {
		Tag              string `yaml:"TAG"`
		Verbosity        string `yaml:"VERBOSITY"`
		Bootnodes        string `yaml:"BOOTNODES"`
		P2PPrivKEY       string `yaml:"P2P_PRIVKEY"`
		ExternalIP       string `yaml:"EXTERNAL_IP"`
		P2PHostDNS       string `yaml:"P2PHostDNS"`
		RPCHost          string `yaml:"RPCHost"`
		RPCPort          string `yaml:"RPCPort"`
		UDPPort          string `yaml:"UDPPort"`
		TCPPort          string `yaml:"TCPPort"`
		MonitoringHost   string `yaml:"MONITORING_HOST"`
		HTTPWeb3Provider string `yaml:"HTTP_WEB3_PROVIDER"`
	}

	Validator struct {
		Tag       string `yaml:"TAG"`
		Verbosity string `yaml:"VERBOSITY"`
	}

	WalletDir          string `yaml:"WALLET_DIR"`
	WalletPasswordFile string `yaml:"WALLET_PASSWORD_FILE"`
}

func LoadDefaults(LuksoSettings *LuksoValues) {
	LuksoSettings.Force = false
	LuksoSettings.Network = "l15"
	LuksoSettings.Orchestrator.Verbosity = "info"
	LuksoSettings.Pandora.Verbosity = "info"
}
