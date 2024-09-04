// func 函数名称(接收的参数 参数的类型，参数2 类型2)

// func函数名称(参数1，参数2，参数3类型）（返回值1，返回值2 类型）{代码块/函数体}

// func函数名称（参数1，参数2，参数3类型）{代码块/函数体}

// func函数名称（参数1，参数2，参数3类型）类型{代码块/函数体}
package main

import (
	"fmt"
	"strings"
)

// 求和函数，返回值有名字，直接return返回
func qiuhe(a, b int) (sum int) {
	sum = a + b
	return
}

// 传过来参数长度不定
func sum(args ...string) string {
	// 接收到的是切片
	fmt.Println("args:", args)
	fmt.Printf("接收到的数据类型args type:%T\n", args)
	msg := strings.Join(args, "--")
	fmt.Println("拼接之后的字符串:", msg)
	return msg
}

func main112() {
	fmt.Println("具有函数名字")
	res := qiuhe(10, 20)
	fmt.Println("res:", res)

	sum("hello", "world", "go")
}
