package main

import (
	"errors"
	"fmt"
)

func connectDB(adrress string, port int) (string, error) {
	if adrress == "" || port == 0 {
		return "", errors.New("数据库连接失败")
	} else {
		return "数据库连接成功", nil
	}
}

func mainvb() {
	s, err := connectDB("127.0.0.1", 3306)
	if err != nil {
		fmt.Println(err)
		panic(err) // 终止程序运行
	} else {
		fmt.Println(s)
	}
}
