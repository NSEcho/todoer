package cmd

import (
	"strconv"

	"github.com/lateralusd/todoer/db"
	"github.com/lateralusd/todoer/helpers"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove task",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := db.InitDatabase(helpers.GetDBPath())
		for _, ids := range args {
			id, err := strconv.Atoi(ids)
			if err != nil {
				return err
			}
			if err := db.Remove(id); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
