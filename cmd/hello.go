package hello

import (
	"fmt"
	"github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"
	"os"
	"path/filepath"
)

func Task() {
	name := tasks.GetInput("name")
	if name == "" {
		tasks.Panic("name input cannot be empty")
	}

	// Compute greeting message
	greeting := fmt.Sprintf("Hello %s!", name)
	tasks.Info(greeting)

	// Create greeting file at the current working directory
	outFilePath := filepath.Join(tasks.GetWorkingDir(), fmt.Sprintf("greeting.txt"))
	err := os.WriteFile(outFilePath, []byte(greeting), 0666)
	if err != nil {
		tasks.Panic(fmt.Sprintf("unable to write to file %s with error: %s", outFilePath, err))
	}

	// Export environment variable containing path to greeting file
	err = tasks.ExportEnvironmentVariable("GREETING_FILE", outFilePath)
	if err != nil {
		tasks.Panic(fmt.Sprintf("failed to set a variable export: %s", err))
	}

	// Save the file's path to task state
	tasks.SetState("pathToFile", outFilePath)

	// Set greeting message as task output
	tasks.SetOutput("greeting", greeting)
}

func Cleanup() {
	// remove greeting file if exists
	outFilePath := tasks.GetState("pathToFile")

	if outFilePath != "" {
		tasks.Info(fmt.Sprintf("cleaning up file at path: %s", outFilePath))
		err := os.Remove(outFilePath)
		if err != nil {
			if !os.IsNotExist(err) {
				tasks.Panic(fmt.Sprintf("unable to remove greeting file: %s", err))
			} else {
				tasks.Info("greeting file not found. skipping...")
			}
		}

	} else {
		tasks.Info("no greeting file recorded in state. skipping...")
	}
}
