package main

import (
	"fmt"
	"reflect"
)

func main7() {
	// 关系运算符
	// == 等于
	//!= 不等于
	// > 大于
	// < 小于
	// >= 大于等于
	// <= 小于等于
	a := 10
	b := 20
	// 是否相等, =和==的区别, =是赋值运算符, ==是比较运算符
	fmt.Println("a==b:", a == b) // false
	fmt.Println("a!=b:", a != b) // true

	// 字符串比较
	s1 := "hello"
	s2 := "world"
	fmt.Println("s1==s2:", s1 == s2) // false , 字符串比较是比较两个字符串的内存地址是否相同, 而不是比较字符串的内容 ,通过ascii码比较

	// 逻辑运算符
	// && 逻辑与
	// || 逻辑或
	//! 逻辑非
	n1 := 1
	n2 := 2
	n3 := 3
	fmt.Println("n1 && n2 && n3:", n1 == 1 && n2 == 2 && n3 == 3) // true

	// 判断字符串是否为字符串类型
	fmt.Println(n1 == n2 || reflect.TypeOf(n3).Kind() == reflect.String)
}
