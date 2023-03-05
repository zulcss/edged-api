package cmd

import "github.com/spf13/cobra"

var nodeCmd = &cobra.Command{
	Use:	"node",
	Short:	"Check the status of the server",
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
