package main

import (
	"fmt"
	"time"
)

// 1. Sample routine
/* func hello() {
	fmt.Println("Hello GoRoutines!!")
} */

// 2. mutliple routines

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'u'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	// 1.
	/* go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Main program!") */

	// 2.
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated!")
}
