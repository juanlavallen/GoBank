package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/juanlavallen/CLI_AppGO"
)

const file = ".todos.json"

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(file); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Example Todo")
		err := todos.Store(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	default:
		fmt.Fprintln(os.Stderr, "Command not valid")
		os.Exit(0)
	}
}
