package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	type student_json struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Class int    `json:"class"`
	}
	id1 := student_json{
		"xie", 18, 1,
	}
	id2, _ := json.Marshal(id1)
	fmt.Print(id1)
	w.Write(id2)

}
func main() {
	port := ":9000"
	log.Println("端口启动成功，端口号为", port)
	go http.HandleFunc("/", hello)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Print("发生了一个错误，错误信息为:", err)
		//log.Panic("错误，端口启动失败错误")
	}

}
