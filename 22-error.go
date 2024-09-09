package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (res int, err error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	res = a / b
	return res, nil
}

func mainzx() {
	// 自定义error
	error := errors.New("这是一个新的错误")
	fmt.Println(error)
	fmt.Println("================")
	res, err := divide(1, 1)
	if err != nil {
		fmt.Println("计算出错", err.Error())
	} else {
		fmt.Println("计算结果为:", res)
	}
}
