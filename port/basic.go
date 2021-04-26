package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type data struct {
	Key string `json:"key"`
	Num int    `json:"num"`
}

func get(urls string) {

	gets, err := http.Get(urls)
	if err != nil {
		print("错误信息为:", err)
		panic("发生了内存错误，读取不到对象，")
	}
	content, err := ioutil.ReadAll(gets.Body)
	if err != nil {
		print("错误信息为:", err)
		panic("发生了内存错误，读取不到对象，错误信息为:")
	}
	fmt.Printf("%s", content)
}
func post(url, contenType string) {
	datas := data{
		Key: "",
		Num: 5,
	}
	datass, err := json.Marshal(datas)
	post, err := http.Post(url, contenType, strings.NewReader(string(datass)))
	if err != nil {
		print("错误信息为:", err)
		panic("发生了内存错误，读取不到对象，")
	}
	content, err := ioutil.ReadAll(post.Body)
	fmt.Printf("%s", content)
}
func tets(url string) {
	content, err := http.Get(url)
	if err != nil {
		panic("报错了，错误内容你自己猜")
	}
	//q:=content.URL
	rs, err := ioutil.ReadAll(content.Body)
	if err != nil {
		panic("炸了")
	}
	fmt.Print(rs)

}
func main() {
	post("https://bduss.loli.fit/pixiv.php", "application/json")
}
