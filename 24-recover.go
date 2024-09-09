package main

import "fmt"

func PrintSlice(s []string) {
	// recover 只能在defer中使用
	defer func() { // 匿名函数,()作用是立即执行
		if err := recover(); err != nil {
			// recover 捕获panic到错误后，程序不会崩溃，会执行defer中的代码
			fmt.Println("捕获到错误：", err)
		}
	}()
	fmt.Println(s)
	// 打印切片的第三个值
	fmt.Println(s[2])
}

func mainqasw12() {
	s := []string{"1", "2"}
	PrintSlice(s)
}
