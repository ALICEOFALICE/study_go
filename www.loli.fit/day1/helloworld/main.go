package main

import "fmt"
import "math"

func main() {
	var a float64
	var b float64
	fmt.Printf("请输入数值：")
	fmt.Scanf("%f", &a)
	if a > 1000 {
		fmt.Printf("检测到数值超出1000，请不要输出超出1000的值")
		return
	}
	b = math.Sqrt(a)
	fmt.Printf("结果为：%.0f", b)
	return
}
