package cmd

import (
	"github.com/automation-co/husky/internal/lib"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install hooks",
	Long:  `Install hooks from the .hooks folder`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Install(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
