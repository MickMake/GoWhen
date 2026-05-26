package main

import (
	"fmt"
	"os"

	"github.com/MickMake/GoUnify/Only"
	"GoWhen/cmd"
)

func main() {
	var err error

	for range Only.Once {
		err = cmd.Execute()
		if err != nil {
			break
		}
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
