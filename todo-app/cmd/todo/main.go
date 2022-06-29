package main

import (
	"flag"
	"fmt"
	"os"

	".../todo"
)

const (
	todoFile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add task")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("sample task")
		err := todos.Store(todoFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "no task")
		os.Exit(0)
	}
}
