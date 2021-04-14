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
func abab(a *int) {
	println(*a)
}
func main() {
	str := new(string)
	*str = "GO语言开发"
	println(*str)
}
