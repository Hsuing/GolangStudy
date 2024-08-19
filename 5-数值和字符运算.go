package main

import "fmt"

func numOperations(a, b int) {
	fmt.Println("a + b=", a+b)
	fmt.Printf("%d - %d=%d\n", a, b, a-b)
	fmt.Printf("%d * %d=%d\n", a, b, a*b)

	// 除法运算结果为浮点型, 因此需要使用float64()进行转换
	fmt.Printf("%d / %d=%f\n", a, b, float64(a)/float64(b))
	fmt.Printf("%d %% %d=%d\n", a, b, a%b)
}

// 字符运算
func charOperations(s1, s2 string) {
	fmt.Printf("s1和s2拼接后=%s\n", s1+s2)
}

func main5() {
	// 数值运算
	a1 := 1
	b1 := 2
	fmt.Println("a1 + b1=", a1+b1)
	numOperations(1, 2)
	charOperations("1", "2")

	// 自增自减运算
	a2 := 1
	a2++
	fmt.Println("a2++=", a2)
	a2--
	fmt.Println("a2--=", a2)
}
