package cmd

import (
	"github.com/lateralusd/todoer/db"
	"github.com/lateralusd/todoer/helpers"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "get tasks in the last 24hs",
	RunE: func(cmd *cobra.Command, args []string) error {
		database := db.InitDatabase(helpers.GetDBPath())
		tasks, err := database.Last24HTasks()
		if err != nil {
			return err
		}

		c, err := cmd.Flags().GetBool("completed")
		if err != nil {
			return err
		}

		i, err := cmd.Flags().GetBool("incompleted")
		if err != nil {
			return err
		}

		if c {
			var completedTasks []db.Task
			for _, task := range tasks {
				if task.Completed {
					completedTasks = append(completedTasks, task)
				}
			}
			helpers.PrintTasks(completedTasks, "Completed tasks for today")
			return nil
		}

		if i {
			var incompletedTasks []db.Task
			for _, task := range tasks {
				if !task.Completed {
					incompletedTasks = append(incompletedTasks, task)
				}
			}
			helpers.PrintTasks(incompletedTasks, "Incompleted tasks for today")
			return nil
		}

		helpers.PrintTasks(tasks, "Tasks for today")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(todayCmd)
	todayCmd.Flags().BoolP("completed", "c", false, "show completed tasks for today")
	todayCmd.Flags().BoolP("incompleted", "i", false, "show uncompleted tasks for today")
}
