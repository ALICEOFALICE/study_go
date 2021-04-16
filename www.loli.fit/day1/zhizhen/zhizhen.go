package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
func flags() {

	flag.Parse()
	fmt.Println(*mode)
}
func abab(a *string) {
	fmt.Println(*a)
}
func main() {
	str := new(string)
	*str = "GO语言开发"
	strs := "123"
	println(*str)
	abab(&strs)

}
