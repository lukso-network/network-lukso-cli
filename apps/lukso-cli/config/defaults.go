package config

type LuksoValues struct {
	Network string `yaml:"NETWORK"`
	Config  string
	OrchestratorTag string `yaml:"ORCHESTRATOR_TAG"`
	PandoraTag string `yaml:"PANDORA_TAG"`
	VanguardTag string `yaml:"VANGUARD_TAG"`
	ValidatorTag string `yaml:"VALIDATOR_TAG"`
	
}

var DefaultValues LuksoValues

func LoadDefaults() LuksoValues {
	DefaultValues.Network = "l15"
	return DefaultValues
}
