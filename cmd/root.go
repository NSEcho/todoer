package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todoer",
	Short: "tasks manager",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
