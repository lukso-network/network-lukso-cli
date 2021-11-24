package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var ConfigValues LuksoValues

func LoadConfig(configFilePath string) error {
	buf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	c := &ConfigValues
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("in file %q: %v", configFilePath, err)
	}

	return nil
}
