package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand_num(ch1 chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 100; i++ {
		ch1 <- int(rand.Int31n(1000))
	}
	close(ch1)
}
func x3(ch1, ch2 chan int) {
	for ret := range ch1 {
		ch2 <- ret * 3
	}
	close(ch2)
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	//创建两个channel，然后初始化他们
	go rand_num(ch1)
	go x3(ch1, ch2)
	for v1 := range ch2 {
		fmt.Printf("输出值为:%d\n", v1)
	}
}
