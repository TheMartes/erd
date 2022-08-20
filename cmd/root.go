package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "erd",
	Short: "ElasticSearch Replication Daemon",
	Long: `Fast, Scalable & Resilient replication daemon from MongoDB to ElasticSearch.
For more information, bug reports & feature request, please visit https://github.com/TheMartes/erd`,
	Run: func(cmd *cobra.Command, args []string) {
		// Empty bcs of help
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
