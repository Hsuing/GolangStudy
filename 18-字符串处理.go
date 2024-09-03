package main

import (
	"fmt"
	"strings"
)

func mainqwer() {
	// 反引号和
	// 字符串处理
	// 1.字符串修剪
	s1 := "  hello world ,"
	fmt.Println("去掉逗号:", strings.Trim(s1, ","))

	// 2.字符串分割
	s2 := "hello,world,go"
	fmt.Println("分割字符串:", strings.Split(s2, ","))

	// 3.字符串连接

	// 4.字符串替换
	s4 := "jenkins"
	fmt.Println("字符串替换:", strings.Replace(s4, "j", "J", 1))
	// 5.字符串查找
	s5 := "jenkins"
	fmt.Println("字符串查找:", strings.Contains(s5, "jen"))
	// 6.字符串统计
	s6 := "jenkins"
	fmt.Println("字符统计:", strings.Count(s6, "j"))
	// 7.字符串截取
	s7 := "jenkins"
	fmt.Println("字符串截取:", s7[2:4])
	// 8.字符串转换
	// 小写变大写
	s8 := "jenkins"
	fmt.Println("字符串转换:", strings.ToUpper(s8))
	// 大写变小写
	s99 := "JENKINS"
	fmt.Println("字符串转换:", strings.ToLower(s99))
	// 9.字符串比较
	s9 := "jenkins"
	s10 := "jenkins"
	fmt.Println("字符串比较:", s9 == s10)
	// 10.字符串格式化
}
