package cmd

import (
	"fmt"
	"os"
	"ucli/api"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "create a new project",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		client := api.NewAPIClient(projectApiURL, token)
		project, err := client.CreateProject(name)
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
