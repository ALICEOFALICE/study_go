package main

import "fmt"

type person struct {
	name string
	ages int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		ages: age,
	}
}
func main() {
	p1 := newPerson("谢", 18)

	fmt.Print(p1.name)
}
