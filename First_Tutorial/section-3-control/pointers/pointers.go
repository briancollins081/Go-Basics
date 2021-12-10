package main

import "fmt"

/* func main() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)
} */
// func setTo10(a *int) {
// 	*a = 10
// }

func setTo10(a *int) {
	a = new(int)
	*a = 10
}

func main() {
	// var b *int
	// fmt.Println(b)
	// fmt.Println(*b)

	// b := new(int)
	// fmt.Println(b)
	// fmt.Println(*b)

	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}
