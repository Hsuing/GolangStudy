package main

import "fmt"

type Serson struct {
	Name string
	Age  int
}

// 方法名和结构体方法名相同
// 值类型：相当于声明了一个新的变量，在方法中修改新变量的值不会影响原有的值
// 指针类型：相当于声明了一个新的变量，在方法中修改新变量的值会影响原有的值
func (s *Serson) Print() {
	fmt.Println("Name:", s.Name, "Age:", s.Age)
}

func Sum(a1, b1 int) int {
	return a1 + b1
}

type long int

func (a long) Add(b long) long {
	return a + b
}

func mainwe() {
	s := Serson{
		Name: "han",
		Age:  18,
	}
	s.Print()
	fmt.Println("---------面向对象---------")
	var s2 long = 10
	r := s2.Add(20)
	fmt.Println(r)
	fmt.Println("---------面向过程---------")
	r2 := Sum(10, 20)
	fmt.Println(r2)
}
