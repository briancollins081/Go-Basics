package main

import "fmt"

/*
	A language that supports first class functions allows functions to be assigned to variables,
	passed as arguments to other functions and returned from other functions.
	Go has support for first class functions.
*/

// 4. User defined function types
/* type add func(a int, b int) int */

// 5. Higher-order functions - takes one or more functions as arguments & returns a function as its result
/* func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
} */

// 6. Returning functions from other functions
/* func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
} */

// 7. Closures
// Closures are anonymous functions that access the variables defined outside the body of the function.

// 8. Every closure is bound to its own surrounding variable.
/* func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
} */

// 9. Practical use of first class functions
// Todo: a program that filters a slice of students based on some criteria
/* type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
} */

// 10. Map functions
func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	// 1. Anonymous Functions
	/* a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a) */

	// 2. Call an anonymous function without assigning it to a vriable
	/* func() {
		fmt.Println("Hello anonymous function, unassigned to a variable")
	}() */

	// 3. Pass args to anonymous func
	/* func(n string) {
		fmt.Println("Welcome,", n)
	}("Gophers") */

	// 4.
	/* var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s) */

	// 5.
	/* f := func(a, b int) int {
		return a + b
	}
	simple(f) */

	// 6.
	/* s := simple()
	fmt.Println(s(60, 7)) */

	// 7.
	/* a := 5
	func() {
		fmt.Println("a =", a)
	}() */

	// 8.
	/* a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))
	fmt.Println(a("Gopher"))
	fmt.Println(b("!")) */

	// 9.
	/* s1 := student{
		firstName: "Julius",
		lastName:  "MacInvy",
		grade:     "A",
		country:   "Kenya",
	}
	s2 := student{
		firstName: "Andere",
		lastName:  "Brian",
		grade:     "B",
		country:   "Kenya",
	}
	s3 := student{
		firstName: "Elizabeth",
		lastName:  "Mugo",
		grade:     "C",
		country:   "Kenya",
	}
	s := []student{s1, s2, s3}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f) */

	// 10. Map functions
	a := []int{2, 5, 7, 90, 45, 100}
	r := iMap(a, func(n int) int {
		return n*5 + 4
	})
	fmt.Println(r)
}
