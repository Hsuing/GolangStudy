package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mainfdq() {
	// func 函数名称(接收的参数 参数的类型，参数2 类型2)

	// func函数名称(参数1，参数2，参数3类型）（返回值1，返回值2 类型）{代码块/函数体}

	// func函数名称（参数1，参数2，参数3类型）{代码块/函数体}

	// func函数名称（参数1，参数2，参数3类型）类型{代码块/函数体}

	res := max(10, 20)
	fmt.Println("res:", res)
}
