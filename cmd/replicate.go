package cmd

import (
	"github.com/spf13/cobra"
	erdInit "github.com/themartes/erd/init"
)

func init() {
	rootCmd.AddCommand(replicateCmd)
}

var replicateCmd = &cobra.Command{
	Use:   "replicate",
	Short: "Start replication daemon",
	Run: func(cmd *cobra.Command, args []string) {
		erdInit.InitReplication(false)
	},
}
