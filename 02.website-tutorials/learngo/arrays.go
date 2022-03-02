package main

import "fmt"

func main() {
	// Initialization
	/* var a [3]int
	a[0] = 13
	a[1] = 25
	a[2] = 55
	fmt.Println(a) */

	// Short hand
	// a := [3]int{13, 25, 55}
	// fmt.Println(a)

	// Default values
	// a := [3]int{13}
	// fmt.Println(a)

	// Ignoring the length of the array
	// a := [...]int{13, 25, 55}
	// fmt.Println(a)

	// arrays are value types
	/* a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println(a)
	fmt.Println(b) */

	// Passed by value to functions
	/* num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function: ", num)
	changeLocal(num)
	fmt.Println("after passing to function: ", num) */

	// Length
	// a := [...]float64{67.7, 89.9, 21, 78}
	// fmt.Println("length of a is ", len(a))

	// Iterating using for
	// a := [...]float64{67.7, 89.9, 21, 78}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%d th element of a is %.3f\n", i, a[i])
	// }

	// Iterating using for range
	/* a := [...]float64{67.7, 89.9, 21, 78}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nthe sum of all elements of a", sum) */

	// Multidimensional Array
	a := [3][2]string{{"lion", "tiger"}, {"cat", "dog"}, {"pigeon", "peacock"}}
	printArray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(b)
}
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("Inside function ", num)
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
