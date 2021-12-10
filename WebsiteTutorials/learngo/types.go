package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false
	fmt.Println("a: ", a, "b: ", b)
	c := a && b
	fmt.Println("c: ", c)
	d := a || b
	fmt.Println("d: ", d)

	// Integer types
	fmt.Println("INTEGER TYPES")
	var e int = 89
	f := 95
	fmt.Println("value of e ", e, "and f is ", f)
	fmt.Printf("type of e is %T, size of e is %d \n", e, unsafe.Sizeof(e))
	fmt.Printf("type of f is %T, size of f is %d \n", f, unsafe.Sizeof(f))
	// Floating types
	fmt.Println("FLOATING TYPES")
	g, h := 5.67, 8.97
	fmt.Printf("type of g %T and type of h %T", g, h)
	sum := g + h
	diff := g - h
	fmt.Println("sum ", sum, "diff ", diff)
	no1, no2 := 56, 89
	fmt.Println("sum ", no1+no2, " diff ", no1-no2)

	// Complex types
	fmt.Println("COMPLEX TYPES")
	comp1 := complex(5, 7)
	comp := 6 + 9i
	compAdd := comp1 + comp
	compCuml := comp1 * comp
	fmt.Println("sum: ", compAdd)
	fmt.Println("product: ", compCuml)

	// Strings
	first := "Brian"
	last := "Andre"
	name := first + " " + last
	fmt.Println("My name is ", name)
}
