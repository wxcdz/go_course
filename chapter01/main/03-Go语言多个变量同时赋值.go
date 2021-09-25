package main

import (
	"fmt"
)

// 定义全局变量
var a1 = 100
var a2 = 200
var a3 = "jack"

// 一次性声明
var (
	a4 = 300
	a5 = 666
	a6 = "lucy"
)

func demo_301() {
	// golang如何一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性声明多个变量的方式2:
	var b1, b2, b3 = 100, "tom", 888
	fmt.Println("b1=", b1, "b2=", b2, "b3=", b3)

	// 一次性声明多个变量的方式3: 同样可以使用类型推导
	c1, c2, c3 := 100, "tome", 888
	fmt.Println("c1=", c1, "c2=", c2, "c3=", c3)

	// 打印全局变量
	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)
	fmt.Println("a4=", a4, "a5=", a5, "a6=", a6)
}

func demo_302() {
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a, b)
}

func demo_303() {
	var a int = 100
	var b int = 200

	b, a = a, b
	fmt.Println(a, b)
}

func main() {
	fmt.Println("===========demo301==========")
	demo_301()
	fmt.Println("===========demo302==========")
	demo_302()
	fmt.Println("===========demo303==========")
	demo_303()
}
