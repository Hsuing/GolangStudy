package main

import (
	"fmt"
	"math"
	"reflect"
)

func Bool() {
	// bool类型: true false
	// 1, 声明变量,如果没有初始化,默认值为false
	var a bool
	fmt.Println(a)
	var b bool = true
	fmt.Println(b)

	// 2,自动推倒类型
	c := false
	fmt.Printf("c: %v\n", c)
	fmt.Printf("c type: %T\n", c)

	// 3, 类型转换
	d := int(1)
	fmt.Println(d)
}

func Float() {
	// 声明变量
	var a float32
	a = 1.0
	fmt.Println(a)

	// 自动推倒类型
	b := 3.14
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b type: %T\n", b)
}

func Strings() {
	// 字符串类型: string
	// 1, 声明变量
	var a string
	a = "hello"
	fmt.Printf("a: %v\n", a)

	// 2, 自动推倒类型
	b := "jenkins"
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b type: %T\n", b)

	// 3, 字符串拼接
	c := "hello" + "world"
	fmt.Printf("c: %v\n", c)

	// 4, 字符串长度
	d := len(c)
	fmt.Printf("d: %v\n", d)
}

func main12321() {
	// 数值类型: int int8 int16 int32 int64 uint
	// int: 正负数 uint: 无符号整数
	defaultIntType := 1
	fmt.Println("默认数值类型:", reflect.TypeOf(defaultIntType))
	// int和操作系统有关，不同系统的int类型可能不同，如32位系统的int32，64位系统的int64
	// 对于不同系统的不同int类型，可以使用unsafe包进行转换
	var int64Type int64 = 1
	fmt.Println("int64类型:", reflect.TypeOf(int64Type))

	var unitNum uint = 1
	fmt.Println("uint类型:", reflect.TypeOf(unitNum))
	fmt.Println("int取值范围:", math.MinInt, math.MaxInt)
	fmt.Println("uint取值范围:", uint(math.MaxUint))

	// 数值类型: float32 float64 ,float64精度更高,工作上面常用
	defaultFloatType := 1.0
	fmt.Println("默认浮点数类型:", reflect.TypeOf(defaultFloatType))
	float32Type := float32(1.0)
	fmt.Println("float32类型:", reflect.TypeOf(float32Type))
	float64Type := float64(1.0)
	fmt.Println("float64类型:", reflect.TypeOf(float64Type))

	fmt.Println("--------------------------")
	Bool()
	fmt.Println("############################")
	Float()

	fmt.Println("##############--string--##############")
	Strings()
}
