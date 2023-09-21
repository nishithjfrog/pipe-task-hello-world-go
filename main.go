package main

import (
	"os"
	hello "pipe-task-hello-world-go/cmd"
)

func main() {
	var arg string
	if len(os.Args) < 2 {
		arg = "task"
	} else {
		arg = os.Args[1]
	}
	// handle panics gracefully
	defer func() {
		recover()

	}()
	// "cleanup" to run cleanup. everything else runs task
	if arg == "cleanup" {
		hello.Cleanup()
	} else {
		hello.Task()
	}
}
