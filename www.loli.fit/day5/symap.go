package main

import (
	"fmt"
	"sync"
)

var m = make(map[int]int)
var jsq3 sync.WaitGroup

func get(key int) int {
	return m[key]
}
func set(key int, value int) {
	m[key] = value
}
func main() {
	for i := 0; i < 20; i++ {
		jsq3.Add(1)
		go func(i int) {
			set(i, i+100)
			fmt.Printf("key%v,value:%v", i, get(i))
			jsq3.Done()
		}(i)
	}
	jsq3.Wait()
}
