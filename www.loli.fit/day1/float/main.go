package main

import "fmt"

func main() {
	var i1 float32 = 615464.0 //定义一个32位浮点数
	var i2 float64 = 615464.0 //定义一个64位浮点数
	var i3 = 45646551.0
	fmt.Printf("浮点型默认的是float64:%T\n", i3)
	fmt.Printf("当然手动定义的除外:%T\n", i1)
	fmt.Printf("不妨输出一次%f\n", i2)

}
