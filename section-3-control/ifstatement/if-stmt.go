package main

import "fmt"

func main() {
	a := 10
	// if a > 5 {
	// 	fmt.Println("a is bigger than 5")
	// } else {
	// 	fmt.Println("a is less than or equal to 5")
	// }
	if b := a / 4; b > 5 {
		fmt.Println("b is bigger than 5", a, b)
	} else {
		fmt.Println("b is less than or equal to 5", a, b)
	}
}
