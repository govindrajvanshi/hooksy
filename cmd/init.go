package cmd

import (
	"github.com/govindrajvanshi/hooksy/internal/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hooksy",
	Long: `
hooksy is a tool to help you manage your git hooks.

For more information, please visit
https://github.com/govindrajvanshi/hooksy
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Init(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
