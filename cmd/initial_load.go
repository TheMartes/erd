package cmd

import (
	"github.com/spf13/cobra"
	erdInit "github.com/themartes/erd/init"
)

func init() {
	rootCmd.AddCommand(initialLoadCmd)
}

var initialLoadCmd = &cobra.Command{
	Use:   "initial-load",
	Short: "Start with initial load for a given daemon",
	Long: `******
!! Use with care, this will wipe all of the existent clusters, caches and queues for given replication
******`,

	Run: func(cmd *cobra.Command, args []string) {
		erdInit.InitReplication(true)
	},
}
