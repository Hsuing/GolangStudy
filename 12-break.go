package main

import "fmt"

func gomain12() {
	// break continues
	// break 终止本次循环
	// continues 终止本次循环,继续下一次循环

	for count := 0; count < 20; count++ {
		if count == 10 {
			fmt.Println("count is 10, break")
			break
		}
		fmt.Println("我现在的数值是:", count)
	}
	// continue
	for count := 0; count < 20; count++ {
		if count == 10 {
			fmt.Println("count is 10, break")
			continue
		}
		fmt.Println("我现在的数值是:", count)
	}
}
