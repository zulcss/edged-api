package cmd

import (
	"os"
	"log"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/zulcss/edged/apiserver/pkg/routes"
	"github.com/zulcss/edged/apiserver/pkg/db"
	"github.com/zulcss/edged/apiserver/pkg/config"
	"github.com/zulcss/edged/apiserver/pkg/constants"

)

var (
	ConfigFile string
)

var rootCmd = &cobra.Command{
	Use:	"edged",
	Short:	"Edged rest-api server",
}

var initCmd = &cobra.Command{
	Use:	"init",
	Short: 	"Initialize edged rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the configuration file
		config.ReadConfig(ConfigFile)
		log.Printf("Initializing server for %s on %s", viper.GetString("host.server"), constants.Port)

		// Initialize the database
		db.InitDatabase()
	},
}

var startCmd = &cobra.Command{
	Use:	"start",
	Short:	"Start edged rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		config.ReadConfig(ConfigFile)

		log.Printf("Running server for %s on %s", viper.GetString("host.server"), constants.Port)
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
	rootCmd.PersistentFlags().StringVar(&ConfigFile, "config", "", "config file to load")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(startCmd)
}
