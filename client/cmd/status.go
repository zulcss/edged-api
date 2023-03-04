package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	
	"github.com/zulcss/edged/client/pkg/client"
)

var statusCmd = &cobra.Command{
	Use:	"status",
	Short:	"Check the status of the server",
	Run: RunStatus,
}

func RunStatus(cmd *cobra.Command, args[]string) {
	c := client.NewClient(Endpoint)
	node, _ := c.Status()
	fmt.Println("Server status is: ", node.Status)
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
