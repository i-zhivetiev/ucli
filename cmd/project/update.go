package project

import (
	"fmt"
	"os"
	"ucli/api"
	"ucli/cmd/storage"
	"ucli/helpers"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [pub_key]",
	Short: "Update project",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		pubKey := args[0]

		body, err := helpers.LoadJSON(jsonFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := api.NewAPIClient(storage.ProjectApiURL, storage.Token)
		updatedProject, err := client.UpdateProject(pubKey, body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(updatedProject)
	},
}

var jsonFile string

func init() {
	ProjectCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVar(
		&jsonFile,
		"json",
		"",
		"JSON file with project data",
	)
}
