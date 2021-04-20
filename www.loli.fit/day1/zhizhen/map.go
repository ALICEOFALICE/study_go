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
	fmt.Println("如果键存在，那么会返回True")
	value, ok := ccc["sql"]
	fmt.Println(ok)
	fmt.Printf(value)
	for key, val := range ccc {
		fmt.Print(key, val)
		fmt.Printf("\n")
	}
}
