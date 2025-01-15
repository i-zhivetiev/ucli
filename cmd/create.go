package cmd

import (
	"fmt"
	"os"
	"ucli/api"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		body, err := ReadJSONFromStdin()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client := api.NewAPIClient(projectApiURL, token)
		project, err := client.CreateProject(body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(project)
	},
}

func init() {
	projectCmd.AddCommand(createCmd)
}
