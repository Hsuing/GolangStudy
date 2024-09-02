package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	num := 100
	string1 := strconv.Itoa(num)
	fmt.Println("string1:", string1)
	// 获取类型
	fmt.Println("kk", reflect.TypeOf(string1))

	// 字符串转数值
	stirng2 := "100"
	num2, err := strconv.Atoi(stirng2)
	if err != nil {
		fmt.Println("转换失败，错误信息：", err)
	}
	fmt.Println("num2:", num2)
	num3 := 200 + num2
	fmt.Printf("num3: %v\n", num3)
}
