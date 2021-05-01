package main

import "fmt"

type person struct {
	name string
	ages int
}

//方法作用于特定类型的函数
func newPerson(name string, ages int) *person {
	return &person{
		name: name,
		ages: ages,
	}
}
func (p person) get_name() {
	fmt.Print("该用户叫：", p.name)
}
func (p *person) ageup() {
	p.ages = p.ages + 1
}
func (p *person) get_name_zhizhen() {
	fmt.Print("该用户年龄：", p.ages)
}
func main() {
	eihei := newPerson("xie", 18)
	//eihei1 := newPerson("li",18)
	eihei.get_name()         //获取名称
	eihei.ageup()            //年龄+1，用指针调过去
	eihei.get_name_zhizhen() //获取年龄

}
