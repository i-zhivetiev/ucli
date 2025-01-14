package cmd

import (
	"fmt"
	"ucli/api"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [pub_key]",
	Short: "delete a project",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pubKey := args[0]
		client := api.NewAPIClient(projectApiURL, token)
		err := client.DeleteProject(pubKey)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	projectCmd.AddCommand(deleteCmd)
}
