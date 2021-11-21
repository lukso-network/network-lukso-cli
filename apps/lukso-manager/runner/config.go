package runner

import (
	"fmt"
	"io/ioutil"
	"lukso-manager/shared"

	"github.com/go-yaml/yaml"
)

type NetworkConfig struct {
	GENESISTIME       int    `yaml:"GENESIS_TIME"`
	CHAINID           int    `yaml:"CHAIN_ID"`
	NETWORKID         int    `yaml:"NETWORK_ID"`
	FORKCHOICE        int    `yaml:"FORK_CHOICE"`
	PANDORABOOTNODES  string `yaml:"PANDORA_BOOTNODES"`
	VANGUARDBOOTNODES string `yaml:"VANGUARD_BOOTNODES"`
}

func ReadConfig(network string) (*NetworkConfig, error) {
	fileName := shared.NetworkDir + network + "/config/network-config.yaml"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := &NetworkConfig{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", fileName, err)
	}

	return config, nil
}
