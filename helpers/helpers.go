package helpers

import (
	"fmt"
	"path/filepath"

	"github.com/lateralusd/todoer/db"
	"github.com/mitchellh/go-homedir"
)

const (
	dbName = "todoer.db"
)

func GetDBPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, dbName)
}

func PrintTasks(tasks []db.Task) {
	for _, task := range tasks {
		val := "[ ]"
		if task.Completed {
			val = "[x]"
		}
		tAdded := task.TimeAdded.Format("02.01.2006 15:04:05")
		fmt.Printf("%s %d - %s (%s)\n", val, task.ID, task.Value, tAdded)
	}
}
