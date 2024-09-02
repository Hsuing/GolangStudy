package main

import (
	"fmt"
	"time"
)

func wwmain11() {
	// 每10秒打印一次代码
	for {
		timeNow := time.Now().Format("2006-01-02 15:04:05") // 2006-01-02 15:04:05 go诞生时间
		println(timeNow)                                    // 打印当前时间
		println("Hello, world!")
		time.Sleep(3 * time.Second)
	}
}

func for1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//  算出1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forRange() {
	str := "abc"
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	fmt.Println("================")
	// range表示迭代打印每个元素，默认返回2个值：一个是元素的位置，一个是元素的值
	for k, v := range str { // k表示的是元素的位置（0,1,2）,v 表示的是元素对应的值（a,b,c）
		fmt.Printf("str[%d]=%c\n", k, v)
	}
	// 等同于
	for k := range str {
		fmt.Printf("str[%d]=%c\n", k, str[k])
	}
	/// 又等同于
	fmt.Println("=======。。。。。=========")
	for k := range str {
		fmt.Printf("str[%d]=%c\n", k, str[k])
	}
	fmt.Println("========______========")
	//
	// for _, v := range str { // 如果只想取到元素的值，可以用_来占位
	// 	fmt.Printf("%c\n", v)
	// }
}

func mainfor() {
	// wwmain11()
	// for1()
	forRange()
}
