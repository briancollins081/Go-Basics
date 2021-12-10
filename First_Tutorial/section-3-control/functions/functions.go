package main

import "fmt"

func main() {
	// fmt.Println(addNumbers(5, 9))
	// fmt.Println(addNumbers(30, -12))
	// a := addNumbers(19, 34)
	// fmt.Println(a)

	// fmt.Println("\nReturning multiple values")

	// div, reminder := divAndReminder(2, 3)
	// fmt.Println(div, reminder)

	// div, _ = divAndReminder(10, 4)
	// fmt.Println(div)

	// _, reminder = divAndReminder(100, -100)
	// fmt.Println(reminder)

	// divAndReminder(-1, 20)
	a := 1
	arr := [2]int{2, 5}
	s := "hello"
	doubleFail(a, arr, s)
	fmt.Println("In main:", a, arr, s)
}

// func addNumbers(a int, b int) {
// 	fmt.Println(a + b)
// }

// func addNumbers(a int, b int) int {
// 	return (a + b)
// }

// func divAndReminder(a int, b int) (int, int) {
// 	return a / b, a % b
// }

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2

	}
	s = s + s
	fmt.Println("In doubleFail:", a, arr, s)
}
