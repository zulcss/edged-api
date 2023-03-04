package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Site	string
)

var siteCmd = &cobra.Command{
	Use:	"site",
	Short:	"Site commands",
}

// list
var siteListCmd = &cobra.Command{
	Use:	"list",
	Short:	"List sites",
}

// add
var siteAddCmd = &cobra.Command{
	Use:	"add",
	Short:	"Add new site",
}

// remove
var siteRemoveCmd = &cobra.Command{
	Use:	"remove",
	Short:	"Remove site",
}

// udpate
var siteUpdateCmd = &cobra.Command{
	Use:	"update",
	Short:	"Update site",
}

// get
var siteGetCmd = &cobra.Command{
	Use:	"get",
	Short:	"Get Site Info",
}

func init() {
	rootCmd.AddCommand(siteCmd)

	// Site subcommands
	siteCmd.AddCommand(siteListCmd)
	siteCmd.AddCommand(siteAddCmd)
	siteCmd.AddCommand(siteRemoveCmd)
	siteCmd.AddCommand(siteUpdateCmd)
	siteCmd.AddCommand(siteGetCmd)

	// Site flags
	rootCmd.PersistentFlags().StringVarP(&Site, "site", "s", "", "Site to add")	
}
