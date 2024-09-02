package main

import (
	"fmt"
	"unsafe"
)

func main212121() {
	// 深拷贝 和 浅拷贝
	// 深拷贝：拷贝整个结构体，包括结构体内的指针，指针指向的结构体也会被拷贝。（会指向新的内存地址）
	// 浅拷贝：拷贝整个结构体，但只拷贝结构体内的基础类型数据，指针指向的结构体不会被拷贝。（这个不会指向新的内存地址）

	// 深拷贝和浅拷贝的区别在于，深拷贝会拷贝整个结构体，浅拷贝只拷贝基础类型数据。
	// 建议使用深拷贝，因为浅拷贝容易造成数据混乱。

	// 值类型（深拷贝）：int、float、bool、string、struct、array、bool
	// 引用类型(浅拷贝)：slice、map、interface、channel

	s1 := "hello"
	s2 := s1
	fmt.Println("s1和s2的值:", s1, s2)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1 // 浅拷贝
	fmt.Println("slice1和slice2的值:", slice1, slice2)

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	slice2[0] = 88
	fmt.Println("修改后的slice1:", slice1)
	fmt.Println("修改后的slice2:", slice2)

	// 深度拷贝
	// var slice3 []int
	// copy(slice3, slice1)
	// fmt.Println("深度拷贝后的slice3:", slice3)
	slice3 := make([]int, len(slice1), cap(slice1))
	copy(slice3, slice1)
	fmt.Println("深度拷贝后的slice3:", slice3)
	slice3[0] = 99
	fmt.Println("修改后的slice1:", slice1)
	fmt.Println("修改后的slice3:", slice3)

	// 打印内存地址
	fmt.Println("slice1的内存地址:", unsafe.Pointer(&slice1))
	fmt.Println("slice3的内存地址:", unsafe.Pointer(&slice3))
	fmt.Println("slice1内存地址中第一个元素的地址:", unsafe.Pointer(&slice1[0]))
	fmt.Println("slice2内存地址中第一个元素的地址:", unsafe.Pointer(&slice2[0]))
	fmt.Println("slice3内存地址中第一个元素的地址:", unsafe.Pointer(&slice3[0]))
}
