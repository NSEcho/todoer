package cmd

import (
	"fmt"
	"strings"

	"github.com/lateralusd/todoer/db"
	"github.com/lateralusd/todoer/helpers"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := db.InitDatabase(helpers.GetDBPath())
		task := strings.Join(args, " ")
		if err := db.Save(task); err != nil {
			return err
		}
		fmt.Println("[*] Added")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
