package main

import "fmt"

// if condition {
// }
func main() {
	// num := 10
	// if num%2 == 0 {
	// 	fmt.Println("The number ", num, " is even")
	// 	return
	// }
	// fmt.Println("The number ", num, " is odd")

	num := 11
	if num%2 == 0 {
		fmt.Println("The number ", num, " is even")
		return
	} else {
		fmt.Println("The number ", num, " is odd")
	}

	num = 99
	if num <= 50 {
		fmt.Println(num, " is less than or equal to 50")
	} else if num > 50 && num <= 100 {
		fmt.Println(num, " is between 51 and 100")
	} else {
		fmt.Println(num, " is greater than 100")
	}
	// If with assignments
	// if assignment-statement; condition {
	// }
	if num2 := 10; num2%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}
}
