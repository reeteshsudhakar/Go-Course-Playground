package main

import "fmt"

func main() {
	levelSevenOne()

	ree := person{
		firstName: "Reetesh",
		lastName:  "Sudhakar",
		age:       18,
	}

	fmt.Println(ree)
	changeMe(&ree)
	fmt.Println(ree)
}
