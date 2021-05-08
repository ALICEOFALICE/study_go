package main

import (
	"fmt"
)

func main() {
	var ch1 chan int
	ch1 = make(chan int, 1)
	ch1 <- 10
	var x int
	x = <-ch1
	fmt.Print(x)
	fmt.Print(len(ch1))
	fmt.Print(cap(ch1))
	close(ch1)
}
