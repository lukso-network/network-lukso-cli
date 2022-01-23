package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  `Print version/build information`,
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}

	return &command
}

func printVersion() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("BuildDate: %s\n", BuildDate)
	fmt.Printf("GitCommit: %s\n", GitCommit)
	fmt.Printf("GitBranch: %s\n", GitBranch)
}
