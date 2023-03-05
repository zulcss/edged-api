package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zulcss/edged/client/pkg/client"
)

var nodeStatusCmd = &cobra.Command{
	Use: 	"status",
	Short:	"Query rest-api heartbeat",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient(Endpoint)
		node, err := c.Status()
		if err != nil {
			fmt.Println("Failed to conenct to", Endpoint)
			os.Exit(-1)
		}
		fmt.Printf("Server %s is %s\n", Endpoint, node.Status)
	},
}

func init() {
	nodeCmd.AddCommand(nodeStatusCmd)
}
