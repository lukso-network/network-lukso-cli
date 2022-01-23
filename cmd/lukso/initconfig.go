package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	viper.SetConfigFile(flags.ConfigFile)

	viper.AutomaticEnv()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(parseLevel(flags.LogLevel))

	if err := viper.ReadInConfig(); err == nil {
		log.Info().Str("config-file", viper.ConfigFileUsed()).Msg("config file loaded")
	}

}

func parseLevel(level string) zerolog.Level {
	switch level {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
