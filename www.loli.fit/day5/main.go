package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wd sync.WaitGroup
var ccc int = 0

func hello(ccc int) {
	fmt.Print("hello", ccc, "\n")
	wd.Done()
}
func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 1048575; i++ {
		wd.Add(1)
		go hello(i)
	}
	fmt.Print("world")
	wd.Wait()
}
