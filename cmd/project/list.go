package project

import (
	"fmt"
	"os"
	"ucli/api"
	"ucli/cmd/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all projects",
	Long:  "",
	Run: func(c *cobra.Command, args []string) {
		client := api.NewAPIClient(storage.ProjectApiURL, storage.Token)
		projects, err := client.GetProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(projects)
	},
}

func init() {
	ProjectCmd.AddCommand(listCmd)
}
