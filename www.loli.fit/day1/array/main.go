package main

import "fmt"

func main() {
	var test [3]int

	test[1] = 3
	fmt.Print(test)
	var test2 = [...]string{"loli", "shi", "sha", "bi"}
	fmt.Print(test2)
	var test3 = [...]string{0: "loli", 2: "shi", 3: "sha", 4: "bi"}
	fmt.Print(test3)
}
