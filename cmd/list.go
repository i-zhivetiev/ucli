package cmd

import (
	"fmt"
	"ucli/api"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all projects",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewAPIClient(projectApiURL, token)
		projects, err := client.GetProjects()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(projects)
	},
}

func init() {
	projectCmd.AddCommand(listCmd)
}
