package main

import (
	"fmt"
	"reflect"
)

// 将 NewInt定义为int类型
type NewInt int

// 将int取别名IntAlias
type IntAlias = int

func demo1501() {
	// 将a声明为NewInt类型
	var a NewInt

	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)

	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

}

// 定义商标结果
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func demo1502() {
	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	a.FakeBrand.Show()

	// 取a的类型反射对象
	ta := reflect.TypeOf(a)

	// 便利a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)

		// 打印成员的字段名和类型
		fmt.Printf("FieldName： %v，FieldType： %v\n", f.Name, f.Type.Name())
	}
}

func main() {
	fmt.Println("===========demo1501==========")
	demo1501()
	fmt.Println("===========demo1502==========")
	demo1502()
}
