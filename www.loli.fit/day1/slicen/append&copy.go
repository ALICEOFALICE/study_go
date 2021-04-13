package main

import "fmt"

func main() {
	slice1 := make([]int, 5, 10)
	slice1 = []int{1, 3, 5, 7, 9}
	fmt.Print(slice1)
	slice1 = append(slice1, 100)
	fmt.Println(slice1)
	fmt.Print("上面就是append，下面我们来个copy")
	fmt.Print(slice1)
	fmt.Print("我们操作一下数据，copy一个变量")
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	fmt.Print("我们定义一个变量切片slice2，他的值为:", slice2, "\n")
	fmt.Print("我们记忆一下变量切片slice1，他的值为:", slice1, "\n")
	fmt.Print("我们操作一下slice1，把第一个数字改为100", "\n")
	slice1[0] = 10000
	fmt.Print("我们操作完毕后，slice1的值为：", slice1, "\n")
	fmt.Print("但是我们查看slice2的值，他没有任何变化，", slice2, "\n")
	fmt.Print("这是因为copy后，切片就没有引入最初变量了，而是导向了一个新的变量，他就被复制了，原来的值也不会有任何的变化")
	slice3 := append(slice2, 300)
	fmt.Print("小小的问题，如果你用append进行添加元素，其实并不是直接操作原切片，而是将你切片的值给覆盖掉，也就是说，你可以这样写\nslice3 = slice(slice2,1)\n这样写的话，slice2的值其实没有变化\n")
	fmt.Print(slice2)
	fmt.Print("\n我们看看slice3的话，值是变了的\n")
	fmt.Print(slice3)
}
