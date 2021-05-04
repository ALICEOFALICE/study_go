package main

import (
	"fmt"
	"os"
	"time"
)

//级别
//log.Debug
//log.Info
//log.Warning
//log.Error

type loger interface {
	writeLog()
}
type log_Debug struct {
	id   int
	time string
	data string
}
type log_Info struct {
	id   int
	time string
	data string
}
type log_Warning struct {
	id   int
	time string
	data string
}
type log_Error struct {
	id   int
	time string
	data string
}

func (receiver log_Debug) writeLog() {
	times := time.Now()
	_, _ = fmt.Fprintf(os.Stdout, "DeBug:ID:%d\tTime:%s\tdata:%s", receiver.id, times, receiver.data)

}
func (receiver log_Info) writeLog() {
	times := time.Now()
	_, _ = fmt.Fprintf(os.Stdout, "Info:ID:%d\tTime:%s\tdata:%s", receiver.id, times, receiver.data)

}
func (receiver log_Warning) writeLog() {
	times := time.Now()
	_, _ = fmt.Fprintf(os.Stdout, "Warning:ID:%d\tTime:%s\tdata:%s", receiver.id, times, receiver.data)

}
func (receiver log_Error) writeLog() {
	times := time.Now()
	_, _ = fmt.Fprintf(os.Stdout, "Error:ID:%d\tTime:%s\tdata:%s", receiver.id, times, receiver.data)

}

func main() {
	c := log_Debug{
		id:   0,
		time: "2020.3.12",
		data: "这是一个测试消息",
	}
	d := log_Error{
		id:   0,
		time: "2020.3.12",
		data: "这是一个测试消息",
	}
	c.writeLog()
	d.writeLog()
}
