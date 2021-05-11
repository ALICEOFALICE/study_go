package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	jsq2   sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	jsq2.Done()
	rwlock.RUnlock()
}
func write() {
	rwlock.Lock()
	time.Sleep(time.Millisecond * 10)
	jsq2.Done()
	rwlock.Unlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		jsq2.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		jsq2.Add(1)
		go write()
	}
	jsq2.Wait()
	fmt.Print(time.Now().Sub(start))
}
