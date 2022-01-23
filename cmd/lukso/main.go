package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/lukso-network/network-lukso-cli/internal/config"
	"github.com/spf13/cobra"
)

const (
	EXIT_CODE_SUCCESS = 0
	EXIT_CODE_FAIL    = 1
)

var (
	// Version contains the current version.
	Version = "none"
	// BuildDate contains a string with the build date.
	BuildDate = "unknown"
	// GitCommit git commit SHA
	GitCommit = "foo"
	// GitBranch git branch
	GitBranch = "main"

	flags *config.Flags

	rootCmd = &cobra.Command{
		Use:   "lukso",
		Short: "lukso",
		Long:  `lukso is a CLI for the lukso deployment`,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd())
	initFlags()
}

func main() {
	exitCode := EXIT_CODE_SUCCESS

	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("Error while executing commands")
		exitCode = EXIT_CODE_FAIL
	}

	os.Exit(exitCode)
}
