package main

import "fmt"

func main() {
	var max = 10
	for i := 0; i <= max; i++ {
		fizzbuzz(i)
	}

	// throwing away the second return from the function
	n, _ := fmt.Println(helloWorld())
	fmt.Println(n)

	shortDeclaration()
}
