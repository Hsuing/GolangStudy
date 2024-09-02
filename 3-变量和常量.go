package main

import (
	"fmt"
	"reflect"
)

func main3() {
	// 声明变量
	// 第一种方式：使用var关键字声明变量，并初始化变量, 变量类型可以省略
	// = 表示赋值操作
	var name string = "Alice"
	fmt.Printf("name: %s\n", name)

	// 第二种方式：使用var关键字声明变量,间接赋值
	var age int
	age = 20
	fmt.Printf("age: %d\n", age)

	// 或者
	var (
		gender string = "female"
		weight int    = 60
	)
	println("gender:", gender, "weight:", weight)

	// 第三种, 直接 := ,让go编译器自动推导变量类型
	height := 170
	fmt.Printf("height: %d\n", height)

	// 推动变量类型, reflect
	fmt.Println("height变量类型:", reflect.TypeOf(height))

	// 常量 , 常量的值不能被修改,const关键字声明
	const a = 1
	fmt.Printf("a: %d\n", a)
	// 常量多重赋值
	const (
		a1 = 10
		b1 = 20
		c1 = 30
	)
	fmt.Println(a1, b1, c1)
}
