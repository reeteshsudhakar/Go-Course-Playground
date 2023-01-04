package main

import (
	"fmt"
	"math"
	"reflect"
)

// exercise 1
func levelSixOneFoo() int {
	return 42
}

func levelSixOneBar() (int, string) {
	return 42, "hello world"
}

// exercise 2
func levelSixTwoFoo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func levelSixTwoBar(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}

	return total
}

// exercise 3
func levelSixThree() {
	defer func() {
		fmt.Println("This was deferred")
	}()

	fmt.Println("this wasn't deferred ")
}

// exercise 4
type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Println("This person's name is", p.firstName, p.lastName, "and is", p.age, "years old.")
}

// exercise 5
type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area of ", reflect.TypeOf(s), ":", s.area())
}

// exercise 6
func levelSixSix() {
	func() {
		fmt.Println("This is an anonymous function.")
	}()

	fmt.Println("This is not an anonymous function.")
}

// exercise 7
func levelSixSeven() {
	f := func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}

	f()
}

// exercise 8
func levelSixEight() func() {
	return func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}
}

// exercise 9
func levelSixNine(f func() int) {
	fmt.Println("the meaning of life is", f())
}

// exercise 10
func levelSixTen() {
	x := 42
	{
		y := 19
		fmt.Println("y can be printed here:", y)
		fmt.Println("we can print x here too:", x)
	}
	fmt.Println("but we can't print y out here")
}
