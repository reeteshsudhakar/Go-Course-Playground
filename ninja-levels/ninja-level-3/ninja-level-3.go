package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func levelThreeOne() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%d ", i)
	}
}

func levelThreeTwo() {
	i := 65
	for {
		if i > 90 {
			break
		}

		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}

		i++
	}
}

func levelThreeThree() {
	birthYear := 2003
	currYear := 2022
	for birthYear <= currYear {
		fmt.Println(birthYear)
		birthYear++
	}
}

func levelThreeFour() {
	birthYear := 2003
	for {
		if birthYear > 2022 {
			break
		} else {
			fmt.Println(birthYear)
		}

		birthYear++
	}
}

func levelThreeFive() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d ", i%4)
	}
	fmt.Println("")
}

func levelThreeSixAndSeven() {
	fmt.Println("------------------")
	fmt.Println(" START SHELL TEST ")
	fmt.Println("------------------")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare("hi", text) == 0 {
		fmt.Println("hi yourself! how's it going?!")
	} else if strings.Compare("bye", text) == 0 {
		fmt.Println("rude.")
	} else {
		fmt.Println("you could start off by saying hi, sheesh.")
	}

	fmt.Println("------------------")
	fmt.Println("  END SHELL TEST  ")
	fmt.Println("------------------")
}

func levelThreeEight() {
	switch {
	case true:
		fmt.Println("you should expect this to print.")
	case false:
		fmt.Println("this shouldn't print.")
	}
}

func levelThreeNine() {
	fmt.Println("---------------------------")
	fmt.Println(" ENTER YOUR FAVORITE SPORT ")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	favSport := strings.Replace(text, "\n", "", -1)

	switch favSport {
	case "football":
		fmt.Println("THAT'S WHAT WE LIKE TO SEE.")

	case "soccer":
		fmt.Println("MMMMMM... interesting choice.")

	default:
		fmt.Println("sus.")
	}
}
