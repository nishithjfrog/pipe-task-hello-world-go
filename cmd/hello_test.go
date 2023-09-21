package hello

import (
	"github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func setupEnvSh() {
	os.Setenv(tasks.TaskStateDirEnv, filepath.Join(os.TempDir(), "state"))
	os.MkdirAll(os.Getenv(tasks.TaskStateDirEnv), 0777)
	os.Setenv(tasks.CurrentWorkingDirEnv, filepath.Join(os.TempDir(), "work"))
	os.MkdirAll(os.Getenv(tasks.CurrentWorkingDirEnv), 0777)
	os.Setenv(tasks.ScriptExtensionEnv, "sh")
	exportsDir := filepath.Join(os.TempDir(), "exports")
	os.Setenv(tasks.TaskExportsFileEnv, filepath.Join(exportsDir, "exports.env"))
	os.MkdirAll(exportsDir, 0777)
	outputsDir := filepath.Join(os.TempDir(), "outputs")
	os.Setenv(tasks.TaskOutputFileEnv, filepath.Join(outputsDir, "outputs.env"))
	os.MkdirAll(outputsDir, 0777)
}

func setupEnvPS() {
	setupEnvSh()
	os.Setenv(tasks.ScriptExtensionEnv, "ps1")
}

func tearDownEnv() {
	tasks.Info("tearing down env")
	os.RemoveAll(tasks.TaskStateDirEnv)
	os.Unsetenv(tasks.TaskStateDirEnv)
	os.RemoveAll(tasks.CurrentWorkingDirEnv)
	os.Unsetenv(tasks.CurrentWorkingDirEnv)
	os.Unsetenv(tasks.ScriptExtensionEnv)
	os.RemoveAll(tasks.TaskExportsFileEnv)
	os.Unsetenv(tasks.TaskExportsFileEnv)
	os.RemoveAll(tasks.TaskOutputFileEnv)
	os.Unsetenv(tasks.TaskOutputFileEnv)

}

func TestCleanupPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no panic")
		}
	}()
	Cleanup()

}
func TestCleanupStatePanic(t *testing.T) {
	setupEnvSh()
	defer tearDownEnv()
	Cleanup()

}
func TestCleanupNoFile(t *testing.T) {

	testStatePath := filepath.Join(os.Getenv(tasks.TaskStateDirEnv), "test.txt")
	setupEnvSh()
	defer tearDownEnv()

	tasks.SetState("pathToFile", testStatePath)
	Cleanup()

}

func TestCleanup(t *testing.T) {

	testStatePath := filepath.Join(os.Getenv(tasks.TaskStateDirEnv), "test.txt")
	setupEnvSh()
	defer tearDownEnv()

	os.WriteFile(testStatePath, []byte("hello world"), 0777)
	defer os.Remove(testStatePath)
	tasks.SetState("pathToFile", testStatePath)
	Cleanup()

}

func TestTaskNoName(t *testing.T) {
	setupEnvSh()
	defer tearDownEnv()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no panic")
		}
	}()
	Task()
}

func TestTask(t *testing.T) {
	testStatePath := filepath.Join(os.Getenv(tasks.TaskStateDirEnv), "test.txt")
	setupEnvSh()
	defer tearDownEnv()
	os.Setenv("IN_name", "TestTask")
	defer os.Unsetenv("IN_name")
	os.WriteFile(testStatePath, []byte(os.Getenv("IN_name")), 0777)
	defer os.Remove(testStatePath)
	tasks.SetState("pathToFile", testStatePath)

	Task()

	assert.FileExists(t, os.Getenv(tasks.TaskExportsFileEnv))
	exportsFile, err := os.ReadFile(os.Getenv(tasks.TaskExportsFileEnv))
	if err != nil {
		t.Errorf("unable to read task exports file")
	}
	assert.Contains(t, string(exportsFile), tasks.GetWorkingDir())
	outputsFile, err := os.ReadFile(os.Getenv(tasks.TaskOutputFileEnv))
	if err != nil {
		t.Errorf("unable to read task exports file")
	}

	assert.Contains(t, string(outputsFile), "Hello TestTask!")
}
