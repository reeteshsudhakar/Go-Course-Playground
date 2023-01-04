package main

import "fmt"

func main() {
	// exercise set 6
	testFoo := levelSixOneFoo()
	test, bar := levelSixOneBar()
	fmt.Println(testFoo, test, bar)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(levelSixTwoFoo(arr...))
	fmt.Println(levelSixTwoBar(arr))

	levelSixThree()

	p1 := person{
		firstName: "Ree",
		lastName:  "Sudhakar",
		age:       19,
	}

	p1.speak()

	s := square{
		length: 56.42,
	}

	c := circle{
		radius: 5.3,
	}

	info(s)
	info(c)

	levelSixSix()
	levelSixSeven()

	g := levelSixEight()
	g()

	levelSixNine(levelSixOneFoo)
	levelSixTen()

}
