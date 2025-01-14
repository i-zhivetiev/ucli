package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ucli",
	Short: "",
	Long:  "",
}

var token string
var projectApiURL string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func init() {
	envToken := getEnv("UCLI_PROJECT_API_TOKEN", "")
	envProjectApiURL := getEnv("UCLI_PROJECT_API_URL", "https://api.uploadcare.com/projects")

	rootCmd.PersistentFlags().StringVar(
		&token,
		"token",
		envToken,
		"Bearer token for authorization (can be set via UCLI_PROJECT_API_TOKEN env)",
	)

	rootCmd.PersistentFlags().StringVar(
		&projectApiURL,
		"base-url",
		envProjectApiURL,
		"Project API URL (can be set via UCLI_PROJECT_API_URL env)",
	)
}
