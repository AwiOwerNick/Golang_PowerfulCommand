package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Complete(num int) error {
	if num <= 0 || num > len(*l) {
		return fmt.Errorf("Item %d does not exist", num)
	}
	(*l)[num-1].Done = true
	(*l)[num-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Delete(num int) error {
	if num <= 0 || num > len(*l) {
		return fmt.Errorf("Item %d does not exist", num)
	}
	*l = append((*l)[:num-1], (*l)[num:]...)
	return nil
}

func (l *List) Save(fileName string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, js, 0644)
}

func (l *List) Get(fileName string) error {
	js, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, l)
}

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}
		// Нумерация с 1, а не с 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

func (l *List) FullList() {
	for i := range *l {
		fmt.Println("Task: ", (*l)[i].Task)
		fmt.Println("Is done?: ", (*l)[i].Done)
		fmt.Println("Created at: ", (*l)[i].CreatedAt)
		fmt.Println("Completed at: ", (*l)[i].CompletedAt)
		fmt.Println()
	}
}
