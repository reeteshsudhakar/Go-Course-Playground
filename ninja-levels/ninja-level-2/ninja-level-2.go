package main

import "fmt"

func levelTwoOne() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)
}

func levelTwoTwo() {
	a := 42 == 42
	b := 46 <= 42
	c := 38 >= 42
	d := 0 != 42
	e := 82 < 83
	f := 96 > 96
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t\n", a, b, c, d, e, f)
}

func levelTwoThree() {
	const a int = 6
	const b = "James Bond"

	fmt.Println(a)
	fmt.Println(b)
}

func levelTwoFour() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}

func levelTwoFive() {
	raw_string_literal := `this is a "raw string literal"`
	fmt.Println(raw_string_literal)
}

func levelTwoSix() {
	const (
		currYear = iota + 2022
		nextYear = iota + currYear
		nextNextYear
		nextNextNextYear
		nextNextNextNextYear
	)

	fmt.Println(nextYear, nextNextYear, nextNextNextYear,
		nextNextNextNextYear)
}
