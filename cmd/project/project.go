package project

import (
	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "project management",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	// cmd.RootCmd.AddCommand(ProjectCmd)
}
