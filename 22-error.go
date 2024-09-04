package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		fmt.Println("文件读取失败", err.Error())
	} else {
		fmt.Println("文件读取成功", string(f))
	}
}
