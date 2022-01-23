package main

import (
	"github.com/lukso-network/network-lukso-cli/internal/config"
)

func initFlags() {
	flags = config.NewFlags()

	rootCmd.Flags().StringVarP(
		flags.LogLevel,
		"logLevel", "l",
		config.DefaultLogLevel,
		"Specify a log level (info, warn, debug, trace, error)",
	)
	rootCmd.Flags().StringVarP(
		flags.LogFile,
		"logFile", "",
		config.DefaultLogFile,
		"Specify the log file",
	)
}
