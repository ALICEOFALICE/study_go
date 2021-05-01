package main

import (
	"fmt"
	"os"
	"sort"
)

var student_all_info = make(map[int]student_info)

type student_info struct {
	id   int
	Name string
	age  int
}

func newStudent(id int, name string, age int) *student_info {
	return &student_info{
		id:   id,
		Name: name,
		age:  age,
	}
}
func addStudent() {
	var name string
	var id int
	var age int
	fmt.Print("您选择了选项2，开始添加用户")
	fmt.Print("请输入学生ID")
	fmt.Scanf("%d", &id)
	fmt.Print("请输入学生姓名")
	fmt.Scanf("%s", &name)
	fmt.Print("请输入学生年龄")
	fmt.Scanf("%d", &age)
	student_all_info[id] = *newStudent(id, name, age)
}
func out_all_student() {
	var keys []int
	for _, key := range student_all_info {
		keys = append(keys, key.id)
	}
	sort.Ints(keys)
	for id, i := range keys {
		fmt.Printf("\n---------------当前索引ID：%d---------------\n", id)
		fmt.Printf("学生ID：%d\n", student_all_info[i].id)
		fmt.Printf("学生姓名：%s\n", student_all_info[i].Name)
		fmt.Printf("学生年龄：%d\n", student_all_info[i].age)
	}
}
func main() {
	out := 1
	for out == 1 {
		fmt.Print("\n请输入您要执行的操作")
		var types int
		fmt.Scanf("%d", &types)
		//if err !=nil{
		//	log.Printf("发生错误：%s",err)
		//	log.Panic("发生错误")
		//}
		switch types {
		case 1:
			{
				fmt.Print("\n您选择了选项1，开始查询全部用户")
				out_all_student()
			}
		case 2:
			{
				addStudent()
			}
		case 3:
			{
				var id int
				fmt.Print("\n您选择了选项3，请输入您要删除的学生")
				fmt.Scanf("%d", &id)
				if _, ok := student_all_info[id]; ok {
					delete(student_all_info, id)
					fmt.Print("\n删除成功")
				} else {
					fmt.Print("\n您输入了一个不存在的值")
				}
			}
		case 4:
			{
				fmt.Print("\n感谢您使用该系统，正在登出")
				out = 0
			}
		}

	}
	os.Exit(1)
}
