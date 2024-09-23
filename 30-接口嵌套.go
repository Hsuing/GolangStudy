package main

import "fmt"

// 代码复用性
type Animal interface {
	Eat()
	Sleep()
}

type Cat interface {
	Animal
}

type Dog interface {
	Animal
}

// 接口隔离性
type General interface {
	Login()
	LogOut()
}

// 会员属性
type Vip interface {
	toDesk()
}

// 普通用户
type User interface {
	General
}

// 会员用户
type VipUser interface {
	General
	Vip
}

// 定义结构体, 实现接口
type UserTpl struct {
	Name string
}

// 实现接口方法
func (u *UserTpl) Login() {
	fmt.Println(u.Name, "登录")
}

func (u *UserTpl) LogOut() {
	fmt.Println(u.Name, "退出")
}

func (u *UserTpl) toDesk() {
	fmt.Println(u.Name, "去前台")
}

func main() {
	// 接口嵌套优点
	// 1. 接口嵌套可以减少代码的重复，提高代码的复用性
	// 2. 实现接口隔离

	var userInterface User
	var u UserTpl
	u.Name = "张三"
	userInterface = &u
	userInterface.Login()
	// userInterface = &UserTpl{Name: "张三"}
	// userInterface.Login()
	// userInterface.LogOut()
}
