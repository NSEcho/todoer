# todoer

Simple TODO manager using the local gorm database.

# Installation

```bash
$ git clone https://github.com/lateralusd/todoer.git
$ cd todoer && go build
$ ./todoer --help
tasks manager

Usage:
  todoer [command]

Available Commands:
  add         add task
  help        Help about any command
  ls          list all tasks
  mark        mark task completed
  rm          remove task
  today       get tasks in the last 24hs

Flags:
  -h, --help   help for todoer

Use "todoer [command] --help" for more information about a command.
```

# Supported commands

* add - add new task
* ls - list tasks(all, completed or incompleted)
* mark - mark task as done
* rm - delete task from the database
* today - get tasks for today
