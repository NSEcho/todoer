package main

import (
	"fmt"
	"os"

	"github.com/lateralusd/todoer/cmd"
)

func main() {
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %+v\n", err)
		os.Exit(1)
	}
}
