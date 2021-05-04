package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name string `json:name`
	Age  int    `json:age`
	Id   int    `json:id`
}

func main() {
	p1 := user{
		Name: "xie",
		Age:  18,
		Id:   10,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		log.Panic("报错了")
	}
	fmt.Printf("%V", b)
}
