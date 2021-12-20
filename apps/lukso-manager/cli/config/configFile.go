package config

import (
	"fmt"
	"io/ioutil"
	"lukso/apps/lukso-manager/settings"

	"gopkg.in/yaml.v3"
)

var ConfigValues settings.Settings

func LoadConfig(Settings *settings.Settings, configFilePath string) error {
	buf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	c := &Settings
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("in file %q: %v", configFilePath, err)
	}

	return nil
}
