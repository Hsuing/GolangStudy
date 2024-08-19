package main

import (
	"fmt"
	"math"
	"reflect"
)

func main6() {
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
}
