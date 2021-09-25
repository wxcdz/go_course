package main

import (
	"fmt"
	"unsafe"
)

// golang中整数类型使用
func demo601() {
	// int8的范围 -128-127
	var n1 int8 = 127
	fmt.Println("n1=", n1)

	// uint8的范围 0-255
	var n2 uint8 = 255
	fmt.Println("n2=", n2)

	// unsafe.Sizeof() 返回变量占用字节数
	fmt.Printf("n2 的 类型 %T  n2 占用的字节数是 %d ", n1, unsafe.Sizeof(n1))
}

func main() {
	fmt.Println("===========demo601==========")
	demo601()
}
