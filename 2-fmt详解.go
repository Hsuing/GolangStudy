package main

import "fmt"

func main2() {
	// 格式化输出
	// %v 打印值的默认格式
	// %s 打印字符串
	// %d 打印整数
	// %f 打印浮点数
	// %t 打印布尔值
	// %q 打印带双引号的字符串
	// %x 打印十六进制数
	// %X 打印十六进制数（大写）
	// %% 打印%符号
	// 格式化输出的例子
	fmt.Printf("Hello, %s!\n", "world")
	fmt.Printf("Age: %d, Height: %f\n", 25, 1.75)
	fmt.Printf("Is it raining? %t\n", true)
	fmt.Printf("Quote: %q\n", "Don't panic")
	//
	fmt.Print("Hello, ") // 直接打印,不换行
	fmt.Print("Hello, ")

	//%s 打印字符串,这两个的区别在于%s 可以在任意位置,Println只能在最后
	fmt.Printf("传递过来的只有一个参数: %s\n", "hello")
	fmt.Println("传递过来的只有一个参数: hello")

	// Sprintf 可以赋值变量
	str := fmt.Sprintf("传递过来的只有一个参数: %s", "hello")
	fmt.Println(str)
}
