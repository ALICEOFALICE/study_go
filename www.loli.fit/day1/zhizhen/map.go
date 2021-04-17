package main

import "fmt"

func main() {
	ccc := make(map[string]string)
	ccc["nihao"] = "sj"
	ccc["didian"] = "成都"
	ccc["sql"] = "php"
	fmt.Print(ccc["nihao"])
	fmt.Print("\n")
	fmt.Print(ccc["didian"])
	fmt.Print("\n")
	fmt.Print(ccc["sql"])
	fmt.Print("\n")
}
