package main

import "fmt"

func main() {
	for i := 1; i < 9; i++ {
		for s := 1; s < 9; s++ {
			if i < s {
				break
			}
			fmt.Printf("%d X %d = %d \t\t", i, s, i*s)
		}
		fmt.Print("\n")
	}
}
