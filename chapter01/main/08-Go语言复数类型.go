package main

import "fmt"

// 复数类型
func demo801() {
	var x1 complex128 = complex(1, 2)
	var y1 complex128 = complex(1, 2)
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(x1 * y1)
	fmt.Println(real(y1), imag(x1))
}

func main() {
	fmt.Println("===========demo801==========")
	demo801()
}
