package cmd

import (
	"fmt"
	"strings"

	"github.com/lateralusd/todoer/db"
	"github.com/lateralusd/todoer/helpers"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search task containing word",
	RunE: func(cmd *cobra.Command, args []string) error {
		database := db.InitDatabase(helpers.GetDBPath())

		wordCase, err := cmd.Flags().GetBool("case")
		if err != nil {
			return err
		}

		wCase := db.CaseInsensitive
		if wordCase {
			wCase = db.CaseSensitive
		}

		tasks, err := database.Search(strings.Join(args, " "), wCase)
		if err != nil {
			return err
		}

		title := fmt.Sprintf("Results for \"%s\"", strings.Join(args, " "))
		helpers.PrintTasks(tasks, title)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolP("case", "c", false, "turn on case sensitive search")
}
