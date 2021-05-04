package main

import (
	"bufio"
	"fmt"
	"os"
)

func buffio() {
	var s string
	fmt.Print("请输入您要输入的字符，按回车结束输入")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("您输入的字符为%s", s)
}
func main() {
	buffio()
	fmt.Fprint(os.Stdout, "发生了一个错误")
}
