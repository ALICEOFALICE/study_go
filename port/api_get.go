package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://byjh.cduestc.cn/Api/List/index")

	buf := bytes.NewBuffer(make([]byte, 0, 512))
	buf.ReadFrom(resp.Body)
	fmt.Println(len(buf.Bytes()))

	fmt.Println(string(buf.Bytes()))
}
