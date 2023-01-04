package main

/*
helpful video:
https://www.youtube.com/watch?v=LvgVSSpwND8&list=PLldlyNutLUqV0IBRN-PWUwz9mMX_znaEN&index=5
*/

import (
	"fmt"
	"time"
)

func main() {
	// creates 2 separate go routines that will only stop when the ENTER key is hit
	go count("sheep")
	go count("fish")
	fmt.Scanln()
}

func count(thing string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
