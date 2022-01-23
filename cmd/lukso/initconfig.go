package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error().Err(err).Msg("error finding home directory")
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName(".lukso")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Info().Str("config-file", viper.ConfigFileUsed()).Msg("config file loaded")
	}
}
