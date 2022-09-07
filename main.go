package main

import (
	"github.com/MickMake/GoUnify/Only"
	"GoWhen/cmd"
	"fmt"
	"os"
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
	}
}
