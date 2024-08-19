package main

import "fmt"

func hemain13() {
	// 数组: 一组具有相同类型并且长度固定的元素的集合
	// 使用场景: 班级,三维老师
	// var 数组名字 = [数组长度]数据类型{数据1,数据2,数据3,....}
	var teachNameArr = [3]string{"张三", "李四", "王五"}
	fmt.Println(teachNameArr)
	teachAgeArr := [3]int{18, 19, 20}
	fmt.Println("第一位老师名字:", teachNameArr[0])

	fmt.Println("第一位老师年龄:", teachAgeArr[0])

	// 获取单个元素的格式: 数组名称[下标],下标是从0开始的
	// 修改单个元素的格式: 数组名称[下标] = 数据

	teachNameArr[0] = "张123"
	fmt.Println("修改后的第一位老师名字:", teachNameArr[0])

	fmt.Println("数组长度:", len(teachNameArr))
	for i := 0; i < len(teachNameArr); i++ {
		fmt.Printf("第%d位老师名字:%s\n", i+1, teachNameArr[i])
	}

	//range 遍历数组 ,推荐使用
	for k, v := range teachAgeArr {
		fmt.Printf("第%d位老师年龄:%d\n", k+1, v)
	}

	//自动推断长度
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr3))
}
