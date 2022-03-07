package cmd

import (
	"github.com/automation-co/husky/internal/lib"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new hook",
	Long:  `Adds a new hook to husky and installs it.`,
	Args:  cobra.ExactArgs(2),
	Example: `husky add pre-commit "
echo 'woof'
"`,
	Run: func(cmd *cobra.Command, args []string) {

		// ARGS:
		hook := args[0]
		cmdStr := args[1]

		lib.Add(hook, cmdStr)
		lib.Install()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
