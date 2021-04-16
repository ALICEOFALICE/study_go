package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]
	s1 = append(s1[0:2], s1[3:]...)
	fmt.Print(s1)
	fmt.Print(a1)
}
