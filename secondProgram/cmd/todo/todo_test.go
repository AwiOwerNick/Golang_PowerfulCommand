package todo

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := List{}
	taskName := "newtask"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestCompleted(t *testing.T) {
	l := List{}
	taskName := "newtask"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	l := List{}
	l.Add("task1")
	l.Add("task2")
	l.Add("task3")
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Task != "task3" {
		t.Errorf("Expected task3, got %q instead.", l[1].Task)
	}
}

func TestGetSet(t *testing.T) {
	l1 := List{}
	l2 := List{}
	l1.Add("task")
	tempF, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tempF.Name())
	if err := l1.Save(tempF.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tempF.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	if l2[0].Task != "task" {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
