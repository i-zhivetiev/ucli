package project

import (
	"fmt"

	"ucli/api"
	"ucli/cmd/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [pub_key]",
	Short: "delete a project",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		pubKey := args[0]
		client := api.NewAPIClient(storage.ProjectApiURL, storage.Token)
		err := client.DeleteProject(pubKey)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	ProjectCmd.AddCommand(deleteCmd)
}
