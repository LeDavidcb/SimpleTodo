package pkg

import (
	"fmt"
	"time"
)

type Todo struct {
	Date time.Time
	Due  time.Time
	Desc string
}
type TodoList []Todo

func (self *TodoList) Diplay() {
	for _, val := range *self {
		fmt.Printf("date: %v, due: %v\nDesc: %v\n", val.Date, val.Due, val.Desc)
	}
}
