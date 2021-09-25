package main

import "fmt"

// golang中bool类型使用
func main() {
	fmt.Println("===========demo901==========")
	demo901()
	fmt.Println("===========demo902==========")
	fmt.Println(demo_902(false))
	fmt.Println("===========demo903==========")
	fmt.Println(demo_903(2))

}

func demo_901() {
	var b = false
	// bool类型占用存储空间是1个字节
	// bool类型只能取true或者false
	fmt.Println("b=", b)
}

func demo_902(b bool) int {
	if b {
		return 1
	}
	return 0
}

func demo_903(i int) bool {
	return i != 1
}
