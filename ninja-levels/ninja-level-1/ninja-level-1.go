package main

import "fmt"

func levelOneOne() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func levelOneTwo() {
	var x int
	var y string
	var z bool
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Println(x, y, z)
}

var x = 42
var y = "James Bond"
var z = true

func levelOneThree() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}

func levelOneFour() {
	type Test int
	var x Test
	fmt.Printf("%T, %v\n", x, x)
	x = 42
	fmt.Println(x)
}

type hotdog int

var a hotdog
var b int

func levelOneFive() {
	fmt.Printf("%T, %v\n", a, a)
	a = 19
	fmt.Println(a)
	b = int(a)
	fmt.Printf("%T, %v\n", a, a)
	fmt.Println(a, b)
}
