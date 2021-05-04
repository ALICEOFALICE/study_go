package main

import "fmt"

type b struct {
	name string
}
type a struct {
	age string
	b   //只要一个构造体中包含另一个构造体，那么这个构造体将会自动继承另一个构造体可以实现的所有方法
}

func (as a) fcs() {
	fmt.Print("你调用了a专属方法")
}
func (ass b) fce() {
	fmt.Print("你调用了b专属方法")
}
func main() {
	pp := a{
		age: "15",
		b:   b{name: "ss"},
	}
	pp.fce()
	pp.fcs()
}
