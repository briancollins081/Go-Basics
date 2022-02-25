package main

import "fmt"

/*
	A language that supports first class functions allows functions to be assigned to variables,
	passed as arguments to other functions and returned from other functions.
	Go has support for first class functions.
*/

func main() {
	// 1. Anonymous Functions
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
}
