package cmd

import (
	"github.com/automation-co/husky/internal/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize husky",
	Long: `
husky is a tool to help you manage your git hooks.

For more information, please visit
https://github.com/automation-co/husky
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
