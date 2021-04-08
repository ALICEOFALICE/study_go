package main

import "fmt"

func main() {
	age := 17
	if age > 18 {
		fmt.Printf("首家线上赌场开业了")
	} else if age > 16 {
		fmt.Printf("等两年来吧")
	} else {
		fmt.Printf("鬼！！！")
	}
}
