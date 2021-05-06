package main

import "fmt"

type name interface {
}

func main() {
	var ccc name
	ccc = 123
	fmt.Printf("TYPE=%T,VALUE=%v\n", ccc, ccc)
	ccc = "xie"
	fmt.Printf("TYPE=%T,VALUE=%v\n", ccc, ccc)
	ccc = false
	fmt.Printf("TYPE=%T,VALUE=%v\n", ccc, ccc)
}
