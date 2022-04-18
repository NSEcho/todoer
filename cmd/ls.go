package cmd

import (
	"github.com/lateralusd/todoer/db"
	"github.com/lateralusd/todoer/helpers"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := db.InitDatabase(helpers.GetDBPath())
		tasks, err := db.All()
		if err != nil {
			return err
		}

		helpers.PrintTasks(tasks, "All tasks")

		return nil
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
