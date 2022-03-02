package main

import "fmt"

func main() {
	// for initialisation; condition; post {
	// }
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	// break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\nline after break in for loop\n")

	// continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\nline after continue in for loop\n")

	// Nested loops
	n := 5
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Labels
	// 1. Without label
	fmt.Println("Break without label")
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
	fmt.Println()
	// 2. With label
	fmt.Println("Break with label")
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}

	// Omit initialization & incremental
	i := 0
	for i <= 10 { // for ;i <= 10;
		fmt.Printf("%d ", i)
		i += 2
	}

	// Multiple variables
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
