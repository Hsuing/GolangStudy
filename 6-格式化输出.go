package main

import "fmt"

func Shuchu() {
	a := 10
	b := "hello"
	c := 3.14
	d := true
	// %T 操作变量所属类型
	fmt.Printf("a: %T\n,b: %T\n,c: %T\n,d: %T\n", a, b, c, d)

	// %d 对应整形格式输出
	// %s 对应字符串格式输出
	// %c 对应字符格式输出
	// %f 对应浮点格式输出
	fmt.Printf("a: %d\n,b: %s\n,c: %f\n,d: %t\n", a, b, c, d)

	fmt.Println("#########################")

	// %v 自动匹配格式
	fmt.Printf("a: %v\n,b: %v\n,c: %v\n,d: %v\n", a, b, c, d)
}

func Scan2() {
	var name string
	var age int
	fmt.Println("请输入姓名和年龄:")
	// fmt.Scanln(&name, &age)
	fmt.Scan(&name, &age)
	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Types() {
	// 类型别名 关键字是type
	// 给 int定义一个类型别名myint
	type myint int
	var a myint // 等价于 var a int

	fmt.Printf("a: %v, type: %T\n", a, a)
}

func main12w() {
	// Shuchu()
	// Scan2()
	Types()
}
