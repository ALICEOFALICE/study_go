package main

import (
	"fmt"
	"strings"
	"unicode"
)

func chartj(chars *string) (ret int) {
	mj := *chars
	num := 0
	strings := ""
	for ccc, value := range mj {
		_ = ccc
		if unicode.Is(unicode.Han, value) {
			num++
			strings = strings + string(value)
		}
	}
	fmt.Printf(strings)
	return num
} //执行中文查找并返回中文字符数
func englishtj(english *string) (ret map[string]int) {
	yj := *english
	yj1 := strings.Split(yj, " ")
	m2 := make(map[string]int, 10)
	for _, value := range yj1 {
		if _, ok := m2[value]; !ok {
			m2[value] = 1
		} else {
			m2[value] = m2[value] + 1
		}
	}
	return m2
} //英文查找并返回数量
func f1() int {
	return 10
}
func f2(id int) (ret int) {
	rel := id * 2.0
	return rel
}
func main() {
	d := "We present the short story The Romance of a Busy Broker by O Henry The story was originally adapted and recorded by the US Department of State"
	fmt.Print(englishtj(&d))
	fmt.Print(f2(f1()))

}
