package main

import "fmt"

func main15() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("最初的数据：", s)
	// s[0] [1] [2] [3] [4] ---> [:5] [x:x] ，左边界包含，右边界不包含
	fmt.Println("取出前3个元素：", s[0:3])

	// 删除一个元素
	fmt.Println("现在的数据：", s)
	s = s[1:]
	fmt.Println("删除第一个元素后的数据：", s)

	s = s[:len(s)-1]
	length := len(s)
	fmt.Println(length)
	fmt.Println("删除最后一个元素后的数据：", s)
}
