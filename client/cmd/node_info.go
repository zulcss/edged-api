package cmd

import (
	"text/tabwriter"
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zulcss/edged/client/pkg/client"
)

var (
	host bool
)

var nodeInfoCmd = &cobra.Command{
	Use: 	"info",
	Short:	"Show rest-api information",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient(Endpoint)
		node, err := c.Status()
		if err != nil {
			fmt.Println("Failed to conenct to", Endpoint)
			os.Exit(-1)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		if host {
			fmt.Println(node.Hostname)
		} else {
			fmt.Fprintf(w, "\nHostname:\t%s\n", node.Hostname)
			fmt.Fprintf(w, "Version:\t%s\n\n", node.Version)
			w.Flush()
		}
	},
}

func init() {
	nodeCmd.AddCommand(nodeInfoCmd)

	nodeInfoCmd.PersistentFlags().BoolVarP(&host, "host", "", false, "Display hostname")
}