package main

import "fmt"

func main() {
	ccc := new(int)
	*ccc = 123
	cccc := ccc
	fmt.Print(*cccc)
	ddd := make([]int, 5, 10)
	ddd = []int{1, 3, 5, 7, 9}
	fmt.Print(ddd)
}
