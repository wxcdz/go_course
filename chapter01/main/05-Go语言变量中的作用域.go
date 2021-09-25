package main

import "fmt"

func demo501() {
	//声明局部变量 a 和 b 并赋值
	var a int = 3
	var b int = 4
	//声明局部变量 c 并计算 a 和 b 的和
	c := a + b
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

//声明全局变量
var a float32 = 3.14

func demo502() {
	//声明局部变量
	var a int = 3
	fmt.Printf("a = %d\n", a)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

func demo503() {
	//局部变量 a 和 b
	var a int = 3
	var b int = 4
	fmt.Printf("demo503() 函数中 a = %d\n", a)
	fmt.Printf("demo503() 函数中 b = %d\n", b)
	c := sum(a, b)
	fmt.Printf("demo503()函数中 c = %d\n", c)
}

func main() {
	fmt.Println("===========dem501==========")
	demo501()
	fmt.Println("===========demo502==========")
	demo502()
	fmt.Println("===========demo503==========")
	demo503()
}
