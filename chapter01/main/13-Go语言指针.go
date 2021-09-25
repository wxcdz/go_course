package main

import (
	"flag"
	"fmt"
)

func demo_1301() {
	fmt.Println("============demo_1301============")
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p\n", &cat, &str)
}

func demo_1302() {
	fmt.Println("============demo_1302============")
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"

	// 对字符串取地址，ptr类型为*string
	ptr := &house

	// 打印ptr的指针地址
	fmt.Printf("ptr type : %T\n", ptr)

	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
}

func demo_1303() {
	fmt.Println("============demo_1303============")
	x, y := 1, 2

	// 交换变量值
	swap(&x, &y)

	// 输出变量值
	fmt.Println(x, y)
}

func swap(a, b *int) {
	// 取a指针的值，赋给临时变量t
	t := *a

	// 取b指针的值，赋给a指针指向的值
	*a = *b

	// 将a指针的值赋给b指针指向的变量
	*b = t
}

var mode = flag.String("mode", "", "process mode")

func demo_1304() {
	fmt.Println("============demo_1304============")
	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)

}

func demo_1305() {
	fmt.Println("============demo_1305============")

	str := new(string)
	*str = "Go语言教程"

	fmt.Println(*str)

}

func main() {
	fmt.Println("===========demo1301==========")
	demo_1301()
	fmt.Println("===========demo1302==========")
	demo_1302()
	fmt.Println("===========demo1303==========")
	demo_1303()
	fmt.Println("===========demo1304==========")
	demo_1304()
	fmt.Println("===========demo1305==========")
	demo_1305()
}
