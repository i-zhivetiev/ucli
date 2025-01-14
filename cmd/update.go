package cmd

import (
	"encoding/json"
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

		bytes, err := os.ReadFile(jsonFile)
		if err != nil {
			fmt.Println(err)
		}

		var body map[string]interface{}
		if err := json.Unmarshal(bytes, &body); err != nil {
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
	projectCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVar(
		&jsonFile,
		"json",
		"",
		"JSON file with project data",
	)
}
