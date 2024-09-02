package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

/*
	计算圆的面积 , 半径为r, 面积为s

函数名：are
参数：r float64
返回值：s float64
*/
func are(r float64) float64 {
	s := math.Pi * r * r

	// 返回值，返回一个float64类型的值，用于用户访问的返回
	return s
}

func main1qaz() {
	r := 5
	s := are(float64(r)) // 类型转换，将int类型的r转换为float64类型的r
	fmt.Println("圆的面积是：", s)

	// 数值转换字符串 ,strconv包
	i := 100
	stringNum := strconv.Itoa(i) // 将int类型的i转换为string类型的stringNum
	fmt.Println("将数值转换为字符串:", stringNum)
	// 获取类型
	fmt.Println("获取类型:", reflect.TypeOf(stringNum))

	// 字符串转换数值 ,strconv包
	stringNum2 := "100"
	num, err := strconv.Atoi(stringNum2) // 将string类型的stringNum2转换为int类型的num
	if err != nil {
		fmt.Println("转换失败，错误信息：", err)
	}
	fmt.Println("num:", num)
}
