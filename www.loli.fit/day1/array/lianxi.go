package main

import "fmt"

func main() {
	var num int
	var num_array = [...]int{1, 3, 5, 7, 9}
	for i := 0; i < len(num_array); i++ {

		num = num + num_array[i]

	}
	fmt.Print(num)
}
