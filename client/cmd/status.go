package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:	"status",
	Short:	"Check the status of the server",
	Run: RunStatus,
}

func RunStatus(cmd *cobra.Command, args[]string) {
	fmt.Println("test")
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
