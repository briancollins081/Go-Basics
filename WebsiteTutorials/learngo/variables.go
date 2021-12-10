package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	fmt.Println("My age is ", age)
	age = 89
	fmt.Println("My age is ", age)

	// Multiple vraiables declaration
	var width, height = 100, 50
	fmt.Println("width is", width, "height is", height)

	var w, h int
	fmt.Println("width is", w, "height is", h)
	w = 900
	h = 560
	fmt.Println("width is", w, "height is", h)
	// Multiple vars
	var (
		name    = "James"
		uage    = 900
		uheight int
	)
	fmt.Println(name, uage, uheight)
	// Shorthand
	count := 10
	fmt.Println(count)

	// multiple vars vm
	fname, fage := "nevreel", 29
	fmt.Println(fname, fage)

	// Reintializing with multiple variables
	a, b := 20, 30
	fmt.Println("a is ", a, "b is ", b)
	c, b := 50, 60
	fmt.Println("b is ", b, "c is ", c)
	b, c = 80, 90
	fmt.Println("b is ", b, "c is ", c)

	e, f := 145.8, 543.8
	g := math.Min(e, f)
	fmt.Println("Minimum value is", g)

}
