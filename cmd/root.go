package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hooksy",
	Short: "Git hooks made easy!",
	Long: `hooksy is a tool to help you manage your git hooks.

For more information, please visit
https://github.com/govindrajvanshi/hooksy
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
