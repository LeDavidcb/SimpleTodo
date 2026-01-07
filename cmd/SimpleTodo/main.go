package main

import (
	"flag"
	"fmt"

	"github.com/SimpleTodo/pkg"
	"github.com/SimpleTodo/pkg/File"
)

func main() {
	var list pkg.TodoList = make(pkg.TodoList, 0)
	file.LoadData(&list)

	listFlag := flag.Bool("list", false, "List out all the elements")
	flag.Parse()

	if *listFlag {
		list.Diplay()
		return
	}

	val, err := pkg.Ask()
	if err != nil {
		fmt.Println(err)
		return
	}

	list = append(list, val)

	err = file.SaveData(list)

	if err != nil {
		fmt.Println(err)
		return
	}
}
