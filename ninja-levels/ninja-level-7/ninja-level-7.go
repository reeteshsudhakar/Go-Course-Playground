package main

import "fmt"

// exercise 1
func levelSevenOne() {
	x := 42
	fmt.Println(&x)
}

// exercise 2
type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(p *person) {
	(*p).firstName = "Bing"
	(*p).lastName = "Bong"
	(*p).age = 19
}
