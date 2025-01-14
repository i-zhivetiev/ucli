package cmd

import (
	"fmt"
	"os"
	"ucli/api"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [pub_key]",
	Short: "Update project",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pub_key := args[0]

		body, err := LoadJSON(jsonFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := api.NewAPIClient(projectApiURL, token)
		updatedProject, err := client.UpdateProject(pub_key, body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(updatedProject)
	},
}

var jsonFile string

func init() {
	updateCmd.Flags().StringVar(
		&jsonFile,
		"json",
		"",
		"JSON file with project data",
	)

	projectCmd.AddCommand(updateCmd)
}
