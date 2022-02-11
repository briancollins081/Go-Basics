package main

func change(val *int) {
	*val = 55
}

func hello() *int {
	i := 5
	return &i
}

func modifyArray(arr *[3]int) {
	(*arr)[0] = 9000
}

func modifyArray2(arr *[3]int) {
	arr[0] = 89990
}

func modifyArray3(sls []int) {
	sls[0] = 99099
}

func main() {
	// Create a pointer
	/* b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a) */

	// Nill value of a pointer
	/* a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	} */

	//Creating pointers with new keyword
	/* size := new(int)
	fmt.Printf("The size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size) */

	// Dereferencing a pointer
	/* b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b) */

	// Passing pointer to a function
	/* a := 78
	fmt.Println("Value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("Value of a before function call is", a) */

	// Returning a pointer from a function
	/* d := hello()
	fmt.Println("Value of d", *d)
	fmt.Println("Value of d", d) */

	// Modify an Array
	/* a := [3]int{90, 80, 100}
	modifyArray(&a)
	fmt.Println(a) */

	// Passing array pointer to array arg
	/* a := [3]int{90, 80, 100}
	modifyArray2(&a)
	fmt.Println(a) */

	// The correct way to pass an array to a function
	/* a := [3]int{76, 54, 12}
	modifyArray3(a[:]) //pass it as a slice
	fmt.Println(a) */

	// No support for pointer arithmetic
	/* t := 788
	s := &t
	fmt.Println(s + 1)
	fmt.Println(*s + 1) */
}
