package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var num_list chan int
var jsqs sync.WaitGroup
var locks sync.Mutex

func rank_num(syncId int) {
	//fmt.Print("--------读取到一个线程--------\n")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3000; i++ {
		fmt.Print("当前运行线程:", syncId, "\n", "开始运行workerID：", i, "\n", "开启互斥锁\n-------\n")
		locks.Lock()
		lessNum := <-num_list
		//fmt.Print("通道内数值为:",lessNum,"\n")

		num_list <- rand.Intn(50) + lessNum
		//fmt.Print("解除互斥锁","\n")
		locks.Unlock()
	}
	fmt.Print("线程运行完毕，释放线程\n")
	jsqs.Done()
}
func main() {

	num_list = make(chan int, 1000)
	num_list <- 0
	for i := 0; i < 3000; i++ {
		jsqs.Add(1)
		go rank_num(i)

	}
	jsqs.Wait()
	close(num_list)
	num_lists := <-num_list
	fmt.Print(num_lists)

}
