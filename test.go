package main

import "fmt"

// 结构体
type Song struct {
	Name string
	Sex  string
}

// 方法
func (s *Song) Eat(food string) {
	fmt.Printf("%s在吃%s", s.Name, food)
}

func main() {
	s := Song{
		Name: "han",
		Sex:  "男",
	}
	s.Eat("苹果")
}
