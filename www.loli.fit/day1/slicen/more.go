package main

import "fmt"

func main() {
	var ddd = make([]int, 5, 10)
	ddd = []int{1, 3, 5, 100, 888}
	for num, arraynum := range ddd {
		fmt.Printf("这是第%d次循环，数组内数据为%d\n", num, arraynum)
	}
}
