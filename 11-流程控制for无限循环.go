package main

import "time"

func wwmain11() {
	// 每10秒打印一次代码
	for {
		timeNow := time.Now().Format("2006-01-02 15:04:05") //2006-01-02 15:04:05 go诞生时间
		println(timeNow)                                    // 打印当前时间
		println("Hello, world!")
		time.Sleep(3 * time.Second)
	}
}
