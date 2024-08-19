package main

import "fmt"

func printPrice(weather string) {
	defaultPrice := 10
	// if else 语句
	// 雨伞: 10 下雨: 20
	if weather == "sunny" {
		fmt.Println("今天晴天,雨伞价格为是:", defaultPrice)
	} else {
		fmt.Println("今天雨天,雨伞价格为:", defaultPrice+10)
	}
}

func printPriceWithWeather(weather string) {
	defaultPrice := 10
	// if else if 语句
	if weather == "lightRain" {
		// 成立的话就执行这里的语句
		fmt.Println("今天小雨,雨伞价格为是:", defaultPrice)
	} else if weather == "heavyRain" {
		fmt.Println("今天大雨,雨伞价格为:", defaultPrice+5)
	} else {
		fmt.Println("今天晴天,雨伞价格为:", defaultPrice)
	}
}

func main8() {
	weather := "sunny"
	printPrice(weather)

	weather = "rainy"
	printPrice(weather)
	fmt.Println("--------------------------------------")
	weather = "lightRain"
	printPriceWithWeather(weather)
}
