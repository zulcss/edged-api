package cmd

import (
	"github.com/zulcss/edged/client/internal/cli"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:	"status",
	Short:	"Check the status of the server",
	Run: func(cmd *cobra.Command, args []string) {
		cli.NewClient(Endpoint)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
