package project

import (
	"fmt"
	"os"
	"ucli/api"
	"ucli/cmd/storage"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [pub_key]",
	Short: "Get project details",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		pubKey := args[0]
		client := api.NewAPIClient(storage.ProjectApiURL, storage.Token)
		createdProject, err := client.GetSingleProject(pubKey)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(createdProject)
	},
}

func init() {
	ProjectCmd.AddCommand(viewCmd)
}
