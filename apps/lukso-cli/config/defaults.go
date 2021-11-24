package config

type LuksoValues struct {
	Network string `yaml:"NETWORK"`
	Config  string
}

var DefaultValues LuksoValues

func LoadDefaults() {
	DefaultValues.Network = "l15"
}
