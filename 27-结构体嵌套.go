package main

import (
	"fmt"
	"reflect"
)

type Person1 struct {
	Name   string
	Age    int
	Mobile Phone // 嵌套结构体
}

type Phone struct {
	Model string
	Price float32
}

type GetInfos struct {
	Person1
	Phone
}

func mainaqwsx() {
	var p Person1
	p.Name = "han"
	p.Age = 20
	p.Mobile.Model = "华为"
	p.Mobile.Price = 1000.00
	fmt.Println(p)
	fmt.Printf("手机型号：%s，价格：%f\n", p.Mobile.Model, p.Mobile.Price)
	fmt.Println("-----------------------------------")
	var m Phone = Phone{"小米", 20000.00}
	var pp Person1
	pp.Name = "张三"
	pp.Age = 20
	pp.Mobile = m
	fmt.Println(pp)
	fmt.Printf("手机型号：%s，价格：%f\n", pp.Mobile.Model, pp.Mobile.Price)
	fmt.Println("pp.Mobile.Price Type:", reflect.TypeOf(pp.Mobile.Price))
}
