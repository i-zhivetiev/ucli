package project

import (
	"fmt"
	"os"

	"ucli/api"
	"ucli/cmd/storage"
	"ucli/helpers"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new project",
	Long:  "",
	Run: func(c *cobra.Command, args []string) {
		body, err := helpers.LoadJSON(jsonFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := api.NewAPIClient(storage.ProjectApiURL, storage.Token)
		project, err := client.CreateProject(body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(project)
	},
}

func init() {
	ProjectCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(
		&jsonFile,
		"json",
		"",
		"JSON file with project data",
	)
}
