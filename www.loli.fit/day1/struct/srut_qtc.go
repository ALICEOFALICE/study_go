package main

import "fmt"

type students struct {
	name   string
	age    int
	nation string
}
type info struct {
	id      int
	student students
}

func main() {
	xie := info{
		id: 1,
		student: students{
			name:   "xie",
			age:    18,
			nation: "china",
		},
	}
	fmt.Print(xie.id)
	fmt.Print(xie.student.name)
}
