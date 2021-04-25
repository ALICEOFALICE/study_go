package main

import "fmt"

func A() {
	defer func() {
		fmt.Print("释放连接")
		if recover() == "error" {
			fmt.Print("栈溢出")
		}
	}()

	panic("error")
	fmt.Print("world")

} //如果panic给的错误的时候，同时在defer中包含了recover（即为panic的错误被recover处理，那么这个报错将会被忽视，程序继续执行）
func B() {
	defer func() {
		fmt.Print("释放连接")
	}()

	panic("error")
	fmt.Print("world")

} //panic打断程序
func main() {
	A()
}
