package main

import (
	"fmt"
	"math"
)

func demo_1401() {
	const (
		e  = 2.7182818
		pi = 3.1415926
	)

	fmt.Println(e, pi)
}

func demo_1402() {
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Monday)
}

func demo_1403() {
	var x1 float32 = math.Pi
	var y1 float64 = math.Pi
	var z1 complex128 = math.Pi
	fmt.Println(x1, y1, z1)

	const Pi64 float64 = math.Pi
	var x2 float32 = float32(Pi64)
	var y2 float64 = Pi64
	var z2 complex128 = complex128(Pi64)
	fmt.Println(x2, y2, z2)
}

func main() {
	fmt.Println("===========demo1401==========")
	demo_1401()
	fmt.Println("===========demo1402==========")
	demo_1402()
	fmt.Println("===========demo1403==========")
	demo_1403()
}
