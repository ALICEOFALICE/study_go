package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, reslut chan int) {
	for job := range jobs {
		fmt.Printf("worker%d,start%d\n", id, job)
		reslut <- job * 2
		time.Sleep(50000)
	}

}

var jsq = sync.WaitGroup{}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 3; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
