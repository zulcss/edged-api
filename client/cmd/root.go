package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var (
	Endpoint	string
)


var rootCmd = &cobra.Command{
	Use:	"stx",
	Short:	"edged rest-api client",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Endpoint, "endpoint", "E", "localhost", "Rest-API endpoint")
}
