package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("请输入您要执行的操作")
	var types int
	fmt.Scanf("%d", &types)
	//if err !=nil{
	//	log.Printf("发生错误：%s",err)
	//	log.Panic("发生错误")
	//}
	if types == 1 {
		fmt.Print("您选择了选项1，开始查询全部用户")
	} else if types == 2 {
		fmt.Print("您选择了选项2，开始添加用户")

	} else if types == 3 {
		fmt.Print("您选择了选项3，开始删除用户")

	} else if types == 4 {
		fmt.Print("感谢您使用该系统，正在登出")
		os.Exit(1)
	}

}
