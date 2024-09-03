package main

import (
	"fmt"
	"strconv"
)

func mainko() {
	// 如果是bool类型，那么只能是1,0,true,false,True,False,TRUE,FALSE
	// strconv.Parsexxx ----> 转换  布尔类型 1,true,True,TRUE,0,false,False,FALSE
	boo := "1"
	boo1, _ := strconv.ParseBool(boo)
	fmt.Println("boo1:", boo1)
}
