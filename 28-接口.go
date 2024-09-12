package main

import "fmt"

type DBconfig struct {
	User     string
	Pass     string
	Host     string
	Port     int
	DataName string
}

/*
type 接口名字 interface {
	方法名(参数列表) 返回值列表
}
*/
// 定义接口,接口中有多少方法,必须全部实现,否则会报错
type DBcommonInterface interface {
	// 方法名 参数列表 返回值列表, 参数列表和返回值列表可以为空,书写方法的时候和这里保持一致
	Insert(string) error
	Delete(string) error
	Update(string) error
	Query(string) error
}

// 定义一个自定义类型去实现接口
type Mysql struct {
	config DBconfig
}

// 实现接口的方法，这里的方法名和参数列表和返回值列表必须和接口中的方法名、参数列表和返回值列表保持一致
func (m Mysql) Insert(sql string) error {
	fmt.Println("插入数据到Mysql,insert sql:", sql)
	return nil
}

// 定义删除方法
func (m Mysql) Delete(sql string) error {
	fmt.Println("删除数据到Mysql,delete sql:", sql)
	return nil
}

// 定义查询方法
func (m Mysql) Query(sql string) error {
	fmt.Println("查询数据到Mysql,query sql:", sql)
	return nil
}

// 定义更新方法
func (m Mysql) Update(sql string) error {
	fmt.Println("更新数据到Mysql,update sql:", sql)
	return nil
}

func main() {
	db := DBconfig{
		User:     "root",
		Pass:     "123456",
		Host:     "127.0.0.1",
		Port:     3306,
		DataName: "test",
	}
	fmt.Println("config:", db)

	// 声明接口变量
	var dbInterface DBcommonInterface

	// 自定义类型赋值
	var m Mysql
	// m.config = config
	// m.Insert("insert")

	m.config = db
	// 接口赋值
	dbInterface = m
	dbInterface.Insert("insert")
}
