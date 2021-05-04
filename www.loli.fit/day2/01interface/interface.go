package main

import "fmt"

type dog struct{}
type cat struct{}
type man struct{}

func (receiver cat) say() {
	fmt.Print("喵喵喵")
}
func (receiver dog) say() {
	fmt.Print("汪汪汪")
}
func (receiver man) say() {
	fmt.Print("嗷嗷嗷")
}

//定义一个抽象的类型，只要实现了sya()这个方法的类型都可以成为sayer
type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say()
}
func main() {
	var c1 cat
	var d1 dog
	var e1 man
	da(c1)
	da(d1)
	da(e1)
	var ss sayer //定义一个接口类型的变量
	ss = c1
	da(ss)
	fmt.Print(ss)
}
