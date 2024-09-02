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

func IF2() {
	a := 10
	if a > 10 {
		fmt.Println("a>10")
	} else {
		fmt.Println("a<=10")
	}

	if a := 10; a < 10 {
		fmt.Println("a<10")
	} else {
		fmt.Println("a>=10")
	}
	// 注意，上边虽然两个例子里边都使用了变量a，但是并不会报冲突。
	// 因为上边是两个独立的a，而这里是一个a，只不过是if语句里边又重新声明了一个a
}

func main2a() {
	weather := "sunny"
	printPrice(weather)

	weather = "rainy"
	printPrice(weather)
	fmt.Println("--------------------------------------")
	weather = "lightRain"
	printPriceWithWeather(weather)
	fmt.Println("--------------------------------------")
	IF2()
}
