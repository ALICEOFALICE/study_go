package main

import "fmt"

func A() {
	fmt.Print("hello")
	panic("error")
	fmt.Print("world")
}
func main() {
	A()
}
