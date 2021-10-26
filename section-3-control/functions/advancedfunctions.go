package main

import "fmt"

// func addOne(a int) int {
// 	return a + 1
// }

/* func main() {

// Assigning a function

// myAddOne := addOne
// fmt.Println(addOne(1))
// fmt.Println(myAddOne(1))

// Declaring a function inside another function must use this approach

myAddOne := func(a int) int {
	return a + 1
}
fmt.Println(myAddOne(1))

} */

// Passing functions to functions

/* func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func main() {
	printOperation(1, addOne)
	printOperation(1, addTwo)
} */

// Variables in Functions
/* func main() {
	// b := 2
	// myAddOne := func(a int) int {
	// 	return b + a
	// }
	// fmt.Println(myAddOne(1))

	b := 2
	myAddOne := func(a int) {
		b = b + a
	}
	myAddOne(5)
	fmt.Println(b)
} */

// Returning a functiomn from a function
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}
func main() {
	addOne := makeAdder(1)
	// addTwo := makeAdder(2)
	doubleAddOne := makeDoubler(addOne)

	fmt.Println(addOne(1))
	fmt.Println(doubleAddOne(1))
}
