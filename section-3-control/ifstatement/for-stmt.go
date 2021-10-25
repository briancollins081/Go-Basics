package main

import (
	"fmt"
)

func main() {
	// a := 4
	// for
	/* for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		if i == a {
			continue
		}
		fmt.Println(i)
	} */
	// while
	/* b := 0
	for b < 8 {
		fmt.Println(b)
		b = b + 1
	}
	fmt.Println(b) */

	// infinite while
	/* b := 0
	for {
		b = b + 1
		fmt.Println(b)
		if b > 10 {
			break
		}
	} */

	// For Range Loop
	// s := "Hello, world"
	s := "👋 🌏"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
