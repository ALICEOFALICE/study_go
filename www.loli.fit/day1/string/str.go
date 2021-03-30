package main

import "fmt"

func main() {
	//修改字符串内容
	s2 := "青年大学习"
	fmt.Println(s2)
	s3 := []rune(s2)        //首先转换为列表状态
	s3[0] = '老'             //索引并修改
	fmt.Println(string(s3)) //转换回去
	//关于字符串与字符的区别
	s4 := '你'
	s5 := "你"
	fmt.Printf("字符的你：%d 字符串的你：%v\n", s4, s5)
	fmt.Printf("字符的你类型：%T 字符串的你类型：%T\n", s4, s5)
	//关于类型转换
	s6 := 64
	fmt.Printf("当前s6值%d\n", s6)
	fmt.Printf("当前s6的类型是：%T，开始强制转换\n", s6)
	s7 := float64(s6)
	fmt.Printf("当前s7值%f\n", s7)
	fmt.Printf("我们修改类型为float64，并且赋予到s7，当前s7的类型是：%T", s7)
}
