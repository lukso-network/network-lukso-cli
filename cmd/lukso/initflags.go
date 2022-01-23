package main

import (
	"github.com/lukso-network/network-lukso-cli/internal/config"
)

func initFlags() {
	flags = config.NewFlags()

	rootCmd.PersistentFlags().StringVarP(
		&flags.LogLevel,
		"logLevel", "l",
		config.DefaultLogLevel,
		"Specify a log level (info, warn, debug, trace, error)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&flags.LogFile,
		"logFile", "",
		config.DefaultLogFile,
		"Specify the log file",
	)
	rootCmd.PersistentFlags().StringVarP(
		&flags.ConfigFile,
		"config", "c",
		config.DefaultConfigFile,
		"Specify the config file",
	)
}
