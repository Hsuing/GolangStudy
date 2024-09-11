package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	// 字段名大写，表示public，小写表示private
	Name string
	Age  int
	Addr string
}

func mainzxd() {
	var p Person
	// 进行赋值
	p.Name = "han"
	p.Age = 18
	fmt.Println(p)
	fmt.Println("p类型:", reflect.TypeOf(p))

	// 或者用语法糖方法进行赋值
	p2 := Person{
		Name: "han",
		Age:  18,
		Addr: "beijing",
	}
	fmt.Println(p2)

	// 或者使用
	var p3 Person = Person{
		Name: "han",
		Age:  18,
		Addr: "beijing",
	}
	fmt.Println(p3)
}
