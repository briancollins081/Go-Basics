package main

import "fmt"

/* func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
} */

// USING SLICES

/* func find(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find(89, []int{89, 90, 95})
	find(45, []int{56, 67, 45, 90, 109})
	find(78, []int{38, 56, 98})
	find(87, []int{})
} */

// PASSING A SLICE TO A VARIADIC FUNCTION

/* func find(num int, nums ...int) {
	fmt.Printf("The type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			found = true
			fmt.Println(num, "found at index", i, "in", nums)
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	nums := []int{89, 90, 95}
	find(89, nums...)
} */

// GOTCHA
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}
func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
