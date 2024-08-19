package main

import "fmt"

// 定义全局变量(在函数外), 全局变量可以被任何函数访问,
var golbalVar string = "广告"

// 定义一个函数, 打印功能
func printAd() {
	fmt.Println("广告一")
	fmt.Println("广告二")
	fmt.Println("我是全局变量:", golbalVar)

	// 调用main函数, 中的变量, 这里会报错, 因为main函数中的变量, 不能被函数外的变量访问
	// fmt.Println("我是main函数中的变量:", name)
}

func main4() {
	fmt.Println("程序开始运行")
	// 调用函数
	printAd()

	// 作用域
	// 1. 局部变量: 定义在函数内部的变量, 只能在函数内部访问, 外部不能访问
	// 2. 全局变量: 定义在函数外的变量, 可以被任何函数访问
	name := "小明"
	fmt.Println("函数内的变量:", name)
}
