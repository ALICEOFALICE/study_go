package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileObj, err := ioutil.ReadFile("././www.loli.fit/day2/logpark/readme.txt")

	if err != nil {
		panic("发生了错误，找不到对象")
	}
	str := "你好"
	fmt.Print(string(fileObj))
	ioutil.WriteFile("././www.loli.fit/day2/logpark/readme.txt", []byte(str), 0666)
}
