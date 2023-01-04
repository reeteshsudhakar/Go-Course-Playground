package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favoriteFlavors []string
}

func levelFiveOne() {
	p1 := person{
		firstName:       "Reetesh",
		lastName:        "Sudhakar",
		favoriteFlavors: []string{"Vanilla", "Cookies and Cream", "Mint Chocolate Chip"},
	}

	p2 := person{
		firstName:       "Rohith",
		lastName:        "Sudhakar1",
		favoriteFlavors: []string{"Chocolate", "Sorbet", "Mango"},
	}

	fmt.Println("Name:", p1.firstName, p1.lastName)
	fmt.Println("Favorite Flavors: ")
	for _, v := range p1.favoriteFlavors {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println("Name:", p2.firstName, p2.lastName)
	fmt.Println("Favorite Flavors: ")
	for _, v := range p2.favoriteFlavors {
		fmt.Printf("\t%v\n", v)
	}
}

func levelFiveTwo() {
	p1 := person{
		firstName:       "Reetesh",
		lastName:        "Sudhakar",
		favoriteFlavors: []string{"Vanilla", "Cookies and Cream", "Mint Chocolate Chip"},
	}

	p2 := person{
		firstName:       "Rohith",
		lastName:        "Sudhakar1",
		favoriteFlavors: []string{"Chocolate", "Sorbet", "Mango"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func levelFiveThree() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	c1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	c2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(c1.doors)
	fmt.Println(c2.luxury)
}

func levelFiveFour() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Reetesh",
		lastName:  "Sudhakar",
		age:       32,
	}

	fmt.Println(p1.firstName, p1.lastName, "is", p1.age, "years old")
}
