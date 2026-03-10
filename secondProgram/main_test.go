package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binname  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		binname += ".exe"
	}
	build := exec.Command("go", "build", "-o", binname)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binname, err)
		os.Exit(1)
	}
	fmt.Println("Running tests....")
	result := m.Run()
	fmt.Println("Cleaning up...")
	os.Remove(binname)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"
	task2 := "test task number 2"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binname)
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		str, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf(" 1: %s\n 2: %s\n", task, task2)
		if expected != string(str) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(str))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewTasks", func(t *testing.T) {
		task_mng := "banan\nvanan\ncaban"
		cmd := exec.Command(cmdPath, "-add", task_mng)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

}
