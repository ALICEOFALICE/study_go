package main

import "fmt"

func main() {
	var c = 101
	fmt.Printf("%d\n", c)  //十进制用%d表示
	i2 := 077              //自适应8进制
	fmt.Printf("%o\n", i2) //八进制用%d表示
	i3 := 0x1234567
	fmt.Printf("%x\n", i3) //十六进制用%x表示
	var i4 int16 = 115
	fmt.Printf("%T", i4)
}
