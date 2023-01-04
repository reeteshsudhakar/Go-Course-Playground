package main

import (
	"fmt"
	"reflect"
)

func shortDeclaration() {
	x := 42
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	x = 99
	fmt.Println(x)

	y := 124
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(y))

	z := "Bond. James Bond."
	fmt.Println(z)
	fmt.Println(reflect.TypeOf(z))

}
