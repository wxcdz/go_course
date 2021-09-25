package main

import "fmt"

// golang中string类型使用
func demo1001() {
	// string的基本类型使用
	var address string = "北京长城 110 hello world"
	fmt.Println(address)

	// 字符串一旦赋值了，字符串就不能修改了: 在Go中字符串是不可变的
	var str = "hello"
	// str[0] = 'a' 字符串不可变
	fmt.Println(str)

	// 使用反引号定义多行字符串
	var str2 = `
			func main() {
				var b = false
				// bool类型占用存储空间是1个字节
				// bool类型只能取true或者false
				fmt.Println("b=", b)
			}`
	fmt.Println(str2)

	var str3 = "hello" + ",world"
	fmt.Println(str3)
}

func main() {
	fmt.Println("===========demo1001==========")
	demo1001()
}
