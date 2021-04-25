package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	server_link, err := net.Listen("tcp", "127.0.0.1:10086") //创建一个监听，监听一个指定的地址
	if err != nil {
		log.Print("错误")
		log.Panic("错误了，发生了一些问题，错误系信息为:", err)
	} //判断监听是否成功
	for {
		client, err := server_link.Accept() //创建一个监听对象，将会返回两个，一个conn一个error
		if err != nil {
			log.Print("错误")
			log.Panic("错误了，发生了一些问题，错误系信息为:", err)
		} //判断获取数据是否成功
		go func(client net.Conn) {
			//接收数据
			var message = make([]byte, 1024) //创建一个byte类型的数组。
			for {
				num, err := client.Read(message)
				if err != nil || num < 1024 {
					break
				} //判断是否是一个正常的http包

				fmt.Print(string(message[0:num])) //读取数据

			}
			//返回数据
			client.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type:text/html;charset=utf-8\r\ncookie:wdnmd\r\n\n你好"))
			defer func() {
				client.Close()
			}()
		}(client)
	}
}
