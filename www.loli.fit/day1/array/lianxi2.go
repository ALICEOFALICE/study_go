package main

import "fmt"

func main() {
	nums_array := [...]int{1, 3, 5, 7, 8}
	num := 8
	//for range1,nums := range nums_array{
	//	for range2,nums2 := range nums_array{
	//		if nums==nums2{
	//			continue
	//		}
	//		if nums+nums2 == num {
	//			fmt.Printf("检测到和为8的代码，下标分别为%d,%d\n",range1,range2)
	//		}
	//	}
	//}
	//这里不知道为啥会出现重复值
	for i := 0; i < len(nums_array); i++ {
		for j := i + 1; j < len(nums_array); j++ {
			if nums_array[i]+nums_array[j] == num {
				fmt.Printf("检测到和为8的代码，下标分别为%d,%d\n", i, j)
			}
		}
	}
}
