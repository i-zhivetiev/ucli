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
		pubKey := args[0]

		body, err := ReadJSONFromStdin()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := api.NewAPIClient(projectApiURL, token)
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
	projectCmd.AddCommand(updateCmd)
}
