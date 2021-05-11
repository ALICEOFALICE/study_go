package main

import (
	"fmt"
	"sync"
)

var num int = 1
var jsq1 sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		num++
		lock.Unlock()
	}
	jsq1.Done()
}
func main() {
	jsq1.Add(2)
	go add()
	go add()

	jsq1.Wait()
	fmt.Print(num)
}
