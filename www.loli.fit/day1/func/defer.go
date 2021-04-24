package main

import "fmt"

func ab() {
	fmt.Print("值为：")
}
func ca(x int, t func()) {
	t()
	fmt.Print(x)
}
func dd() int {
	fmt.Print("----\n")
	a := 1
	defer ca(a, ab)
	a = 2
	fmt.Print("----\n")
	fmt.Print("----\n")
	return 1
}
func main() {
	dd()
}
