package main

import "fmt"

func updateInt(a int) {
	a = 100
}

func updateIntByPointer(a *int) {
	*a = 100
}

func mainhj() {
	a := 10
	// 每个变量有2层含义，变量的内存和变量的地址

	fmt.Println("a的值:", a)     // a,变量的内存，也就是存在内存当中的内容
	fmt.Println("a的内存地址:", &a) //&a,变量的地址，也就是内存所在内存当中的位置，也叫指针

	// 语法糖声明
	// sp 是一个指针变量，存储的是a的内存地址
	sp := &a
	fmt.Println("sp的值:", sp)
	fmt.Println("sp的内存地址:", &sp)

	// // 声明整数指针
	// var p *int
	// fmt.Println("p的值:", p) // 没有初始化,默认值是nil
	// p = &a
	// fmt.Println("p的值:", p)
	// fmt.Println("p的内存地址:", &p)

	// *p = 666 //*p操作的不是p的内存，是p所指向的内存(就是a)
	// fmt.Println("p的值:", *p)
	updateInt(a)
	fmt.Println("updateInt后a的值:", a)
	updateIntByPointer(&a)
	fmt.Println("真正修改后a的值:", a)
}
