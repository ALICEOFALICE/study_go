package main

import "fmt"

func main() {
	var slic []int
	slic = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Print(slic)
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Print(array[0:4])
	fmt.Print(array[3:])
	fmt.Print(array[:])
	fmt.Print("\n")
	v2 := slic[:5]
	fmt.Print("数组v2的切片为:", v2, "\n")
	fmt.Printf("他的长度为:%d,他的容量为:%d", len(v2), cap(v2))
	fmt.Print("\n但是容量是向后取长的，也就是说，如果切片的起点不是数组的起点，那么他的容量将会变小，比如\n")
	v3 := slic[2:5]
	fmt.Printf("比如V3，他的长度为:%d,他的容量为:%d.\n这里的容量就被压缩了，因为他的取值范围是[2:5]，在索引2前面的就被抛弃了，只有后面的还在\n", len(v3), cap(v3))
	fmt.Print("然而，神奇的是，当我们尝试修改一个被二次赋予的切片（切片套切片的时候）原来的值也会变，比如我们马上修改V3（slic[2:5]）的第一个值为10，然后在输出，结果是\n")
	v3[1] = 10
	fmt.Print(slic)
	fmt.Print("\n这说明赋予的切片是引用的原来的数组，当后面的值修改了，之前的也会被修改")
}
