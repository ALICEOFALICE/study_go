package main

import "fmt"

func test1(x, y int) int {
	return (x * y * x * y)
}
func test2(function1 func(x, y int) int, x1, y1 int) func() {
	test2_2 := func() {
		function1(x1, y1)

	}
	return test2_2
}
func test3(x int) func(int) int {

	//fmt.Print(1)
	return func(y int) int {
		//fmt.Print(12)
		return (x + y + x + y)
	}
}
func student(a int) func(int) int {

	var c = 0 + a
	return func(b int) int {
		c = c + b
		return c
	}

}
func main() {
	//fmt.Print(func(x,y int)int{
	//	return (x*y)
	//}(1,3))//这就是一个立即执行函数
	//fmt.Printf("%T",test2(test1,1,2))//这里的执行特别复杂，首先我们导入了一个函数，然后倒入x1，y1并且在函数中创建一个匿名函数，这匿名函数内会执行test1这个函数，并且返回一个单func()，这时候将变量返回就可以直接将test2作为一个变量赋函数使用
	//test3_1 := test3(100)
	//fmt.Print(test3_1(200))
	ccc := student(12)
	lists := [...]int{1, 2, 3, 4, 5}
	for _, list := range lists {
		fmt.Println(ccc(list))
	} //这样做可以让c的值被保存，而不至于函数执行结束就被回收，可看到,c的值被直接保存，在下一次相加的时候依旧被继承
}
