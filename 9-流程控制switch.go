package main

import "fmt"

func printPriceSwitch(weather string) {
	defalutPrice := 10
	switch weather {
	case "Sunny":
		fmt.Println("今天是晴天,雨伞价格为:", defalutPrice)
	case "lightRain":
		fmt.Println("今天是雨天,雨伞价格为:", defalutPrice+5)
	case "heavyRain", "storm":
		fmt.Println("今天是大雨天,雨伞价格为:", defalutPrice+10)
	default:
		fmt.Println("今天是不常见的天气,我不知道雨伞价格,请问客服?")
	}
}

func main9() {
	printPriceSwitch("")
}
