package main

import "fmt"

type name interface {
}

func main() {
	//var ccc name
	//ccc = 123
	//fmt.Printf("TYPE=%T,VALUE=%v\n", ccc, ccc)
	//ccc = "xie"
	//fmt.Printf("TYPE=%T,VALUE=%v\n", ccc, ccc)
	//ccc = false
	//fmt.Printf("TYPE=%T,VALUE=%v\n", ccc, ccc)
	var lists = [...]int{1, 3, 5, 7, 8}
	var c int = 0
	for _, d := range lists {
		c = c + d
	}
	fmt.Print(c)
	var list = [...]int{1, 3, 5, 7, 8}
	for v, _ := range list {
		val := list[v-1] + list[v]
		if val == 8 {
			fmt.Print("当前值为", list[v], "和", list[v+1])
		}
	}
}
