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

import "fmt"

// 1. Panic example
func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	firstName := "Andere"
	fullName(&firstName, nil)
	fmt.Println("return normally from main")
}
