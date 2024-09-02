package main

import "fmt"

func heloomain14() {
	// 切片的长度是不固定的,并且可以动态增长
	// var 切片名称 []切片类型
	var s1 []int
	fmt.Println("最初的切片数据:", s1)

	// 默认的两个属性,一个切片的长度,表示这个切片中有多少个元素
	// 容量: 表示这个切片可以放入多少个元素,当切片的长度超过了容量时,会分配新的内存空间
	// 容量的初始值为0,当向切片中添加元素时,如果容量不足,会自动扩容
	fmt.Println("切片的默认长度:", len(s1))
	fmt.Println("切片的容量:", cap(s1))
	s1 = append(s1, 11121, 989899)
	fmt.Println("长度是:", len(s1))
	fmt.Println("容量是:", cap(s1))

	// 第二种方式创建切片(不常用),使用make函数,指定切片的长度和容量
	s2 := make([]int, 5, 10) // 长度为5,容量为10, 如果是string类型切片,默认的数据是什么
	fmt.Println("切片的默认长度:", len(s2))
	fmt.Println("切片的容量:", cap(s2))
	fmt.Println("最初的切片数据:", s2)
	//
	s2 = append(s2, 1, 2, 3, 4, 5, 6)
	fmt.Println("最初的切片数据:", s2)
	// 现在的切片的容量和长度是多少?
	// 容量*2, 因为切片的容量是可以动态扩容的,当切片的长度超过了容量时,会分配新的内存空间
	fmt.Println("长度:", len(s2))
	fmt.Println("容量:", cap(s2))
	// 修改
	s2[0] = 88
	fmt.Println("修改后的切片数据:", s2)
	// 切片的遍历
	for k, v := range s2 {
		fmt.Printf("第%d个数据是:%d\n", k+1, v)
	}
}
