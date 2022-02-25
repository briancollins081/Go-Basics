/*
 * But there are some situations where the program cannot continue execution after an abnormal condition.
 * In this case, we use panic to prematurely terminate the program.
 * When a function encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to its caller.
 * This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message,
 * followed by the stack trace and then terminates.
 *
 *
 * One important factor is that you should avoid panic and recover and use errors where ever possible.
 * Only in cases where the program just cannot continue execution should panic and recover mechanism be used.
 */

package main

import (
	"fmt"
)

// 1. Panic example
/* func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
} */

// 2. Slice panic
/* func slicePanic() {
	n := []int{5, 7, 6}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")

} */

// 3. Panic behaviour with defer
/* func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
} */

// 4. Recovering from panic
// Note: If recover is called outside the deferred function, it will not stop a panicking sequence
/* func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
} */

// 5. Recover from panic example 2
/* func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack() // get stack trace
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()
	n := []int{7, 5, 3}
	fmt.Println(n[4])
	fmt.Println("normally retruned from invalidSliceAccess method")
} */

// 6. Panic, Recover and Goroutines
/*
	Recover works only when it is called from the same goroutine which is panicking.
	It's not possible to recover from a panic that has happened in a different goroutine.
*/
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a, b int, done chan bool) {
	fmt.Printf("%d ÷ %d = %d", a, b, a/b)
	done <- true

}
func main() {
	/* firstName := "Andere"
	fullName(&firstName, nil)
	fmt.Println("return normally from main") */
	// 2.
	/* slicePanic()
	fmt.Println("normally returned from main") */

	// 3.
	/* defer fmt.Println("deferred call in main")
	lastName := "Collins"
	fullName(nil, &lastName)
	fmt.Println("returned normally from main") */

	// 4.
	/* defer fmt.Println("deferred call in main")
	lastName := "Brian"
	fullName(nil, &lastName)
	fmt.Println("returned normally from main") */

	// 5.
	/* invalidSliceAccess()
	fmt.Println("normally returned from main") */

	// 6.
	sum(5, 0)
	fmt.Println("normally returned from main")
}
