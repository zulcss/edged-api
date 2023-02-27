package cmd

import (
	"os"
	
	"github.com/spf13/cobra"

	"github.com/zulcss/edged/apiserver/pkg/routes"
	"github.com/zulcss/edged/apiserver/pkg/db"
	"github.com/zulcss/edged/apiserver/pkg/config"

)

var rootCmd = &cobra.Command{
	Use:	"edged",
	Short:	"Edged rest-api server",
}

var initCmd = &cobra.Command{
	Use:	"init",
	Short: 	"Initialize edged rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the configuratio file
		config.ReatdConfig()

		// Initialize the database
		db.InitDatabase()
	},
}

var startCmd = &cobra.Command{
	Use:	"start",
	Short:	"Start edged rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		routes.Run()
	},
}

func Execute() {
	err := rootCmd.Execute() 
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(startCmd)
}
