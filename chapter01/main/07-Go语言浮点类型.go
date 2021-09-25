package main

import "fmt"

// golang中浮点类型使用
func demo701() {
	var price float32 = 2222
	fmt.Println("price=", price)
	var num1 float32 = 0.0000012
	var num2 float64 = 64444.333311
	fmt.Println("num1=", num1, "num2=", num2)

	// 尾数部分可能丢失、造成精度损失。 -123.0000901
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	// Golang的浮点型默认声明为float64 类型
	var num5 = 1.1
	fmt.Println("num5的数据类型是 %T", num5)

	// 十进制数形式: 5.12  .512 必须有小数点
	num6 := 5.12
	fmt.Println("num6=", num6)

	// 科学计数法
	num7 := 5.1234e2 // 5.1234 * 10的2次方
	num8 := 5.123e-2 // 5.1234 * 10的-2次方
	fmt.Println("num7=", num7, "num8=", num8)
}

func main() {
	fmt.Println("===========demo701==========")
	demo701()
}
