package main

import "fmt"

type student struct {
	name, id string
}

func f(x *student) {
	x.id = "10000"
}
func main() {
	var R student
	R.id = "10086"
	R.name = "xie"
	f(&R)
	fmt.Print(R.id)
	var c = student{
		name: "li",
		id:   "30",
	}
	fmt.Print(c.name)
}
