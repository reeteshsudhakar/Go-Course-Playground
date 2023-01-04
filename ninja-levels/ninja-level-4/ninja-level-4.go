package main

import "fmt"

func levelFourOne() {
	arr := [5]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", arr)
}

func levelFourTwo() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", arr)
}

func levelFourThree() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := arr[:5]
	b := arr[5:]
	c := arr[2:7]
	d := arr[1:6]
	fmt.Println(a, b, c, d)
}

func levelFourFour() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	arr = append(arr, 52)
	fmt.Println(arr)

	arr = append(arr, 53, 54, 55)
	fmt.Println(arr)

	y := []int{56, 57, 58, 59, 60}
	arr = append(arr, y...)

	fmt.Println(arr)
}

func levelFourFive() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(arr[:3], arr[6:]...)
	fmt.Println(y)
}

func levelFourSix() {
	//arr := make([]string, 0, 50)
	//fmt.Println(arr)
	//fmt.Println(len(arr), cap(arr))
	//
	//arr = append(arr, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	//
	//fmt.Println(len(arr), cap(arr))

	arr := make([]string, 50, 50)

	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))

	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i := 0; i < len(arr); i++ {
		arr[i] = states[i]
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func levelFourSeven() {
	arr := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, v := range arr {
		for _, i := range v {
			fmt.Println(i)
		}
	}
}

func levelFourEight() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Println(k)
		for idx, val := range v {
			fmt.Printf("\t%d: %v\n", idx, val)
		}
	}
}

func levelFourNine() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["Reetesh"] = []string{`Sudhakar`, `Whiskey`, `Computer Science`}

	for k, v := range m {
		fmt.Println(k)
		for idx, val := range v {
			fmt.Printf("\t%d: %v\n", idx, val)
		}
	}
}

func levelFourTen() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(m, `Reetesh`)

	for k, v := range m {
		fmt.Println(k)
		for idx, val := range v {
			fmt.Printf("\t%d: %v\n", idx, val)
		}
	}

}
