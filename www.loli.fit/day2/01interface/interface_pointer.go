package main

import "fmt"

type animal interface {
	move()
	eat(string)
}
type cats struct {
	name string
	feet int8
}

//func (receiver cats) move()  {
//	fmt.Println("走猫步")
//}
//func (receiver cats) eat(food string)  {
//	fmt.Printf("猫吃：%s....\n",food)
//}
func (receiver *cats) move() {
	fmt.Println("走猫步")
}
func (receiver *cats) eat(food string) {
	fmt.Printf("猫吃：%s....\n", food)
}

func main() {
	var a1 animal
	//c1:=cats{
	//	name: "tom",
	//	feet: 4,
	//}//cat
	c2 := &cats{
		name: "假老练",
		feet: 4,
	} //*cat
	//a1=c1//值可以存入a1
	a1 = c2 //指针也可以存进去
	fmt.Println(a1)
}
