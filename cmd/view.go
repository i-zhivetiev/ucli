package cmd

import (
	"fmt"
	"os"
	"ucli/api"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [pub_key]",
	Short: "Get project details",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pubKey := args[0]
		client := api.NewAPIClient(projectApiURL, token)
		createdProject, err := client.GetSingleProject(pubKey)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(createdProject)
	},
}

func init() {
	projectCmd.AddCommand(viewCmd)
}
