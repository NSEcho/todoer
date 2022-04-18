package helpers

import (
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/lateralusd/todoer/db"
	"github.com/mitchellh/go-homedir"
)

const (
	dbName     = "todoer.db"
	timeFormat = "02.01.2006 15:04:05"
)

func GetDBPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, dbName)
}

func PrintTasks(tasks []db.Task, title string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Completed", "Content", "Time Added", "Completed At"})
	for _, task := range tasks {
		val := "[ ]"
		if task.Completed {
			val = "[x]"
		}
		tAdded := task.TimeAdded.Format(timeFormat)
		var tCompl string
		if task.Completed {
			tCompl = task.CompletedAt.Format(timeFormat)
		}
		t.AppendRow(table.Row{task.ID, val, task.Value, tAdded, tCompl})
	}
	t.SetTitle(title)
	t.Style().Title.Align = text.AlignCenter
	t.Render()
}
