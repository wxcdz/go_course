package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}

func demo401() {
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}

func main() {

	fmt.Println("===========demo401==========")
	demo401()
}
