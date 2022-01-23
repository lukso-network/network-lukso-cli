package cluster

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	environment string
)

func NewClusterCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "cluster",
		Short: "",
		Long:  ``,
	}

	command.PersistentFlags().StringVarP(
		&environment,
		"environment", "e",
		"",
		"Specify the environment",
	)
	command.MarkPersistentFlagRequired("environment")

	command.AddCommand(clusterSyncCmd())

	return &command
}

func clusterSyncCmd() *cobra.Command {
	command := cobra.Command{
		Use:   "sync",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			syncCluster(environment)
		},
	}

	return &command
}

func syncCluster(environment string) {
	log.Info().Msgf("Syncing cluster for environment %s", environment)
}
