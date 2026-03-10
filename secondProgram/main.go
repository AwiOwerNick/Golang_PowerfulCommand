package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"secondProgram/cmd/todo"
	"strings"
)

var todoFileName = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Println("This program is a gods plan")
		flag.PrintDefaults()
	}
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	hide := flag.Bool("hide", false, "Hide completed tasks (used with the list flag)")
	fullList := flag.Bool("full", false, "List full info")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Itet to be deleted")
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		if !(*hide) {
			fmt.Print(l)
		} else {
			formatted := ""
			for k, t := range *l {
				prefix := " "
				if t.Done {
					continue
				}
				// Нумерация с 1, а не с 0
				formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
			}
			fmt.Print(formatted)
		}

	case *complete > 0:
		{
			if err := l.Complete(*complete); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	case *add:
		{
			t, err := getTask(os.Stdin, flag.Args()...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			for _, v := range t {
				l.Add(v)
			}

			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	case *delete > 0:
		{
			if err := l.Delete(*delete); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	case *fullList:
		{
			l.FullList()
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) ([]string, error) {
	ret := make([]string, 0, 1)
	if len(args) > 0 {
		ret = append(ret, strings.Join(args, " "))
		return ret, nil
	}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return ret, err
		}
		if len(s.Text()) == 0 {
			break
		}
		ret = append(ret, s.Text())
	}
	if len(ret) == 0 {
		return ret, fmt.Errorf("Task cannot be blank")
	}
	return ret, nil
}
