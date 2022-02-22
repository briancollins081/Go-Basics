package employee

import "fmt"

// 1. Example 1

/* type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, e.TotalLeaves-e.LeavesTaken)
} */

// 2. New Function instead of constructors
/*
 * Go doesn't support constructors.
 * If the zero value of a type is not usable,
 * it is the job of the programmer to unexport the type to prevent access from other packages and
 * also to provide a function named NewT(parameters) which initializes the type T with the required values.
 * It is a convention in Go to name a function that creates a value of type T to NewT(parameters).
 * This will act as a constructor. If the package defines only one type, then it's a convention in Go to name
 * this function just New(parameters) instead of NewT(parameters)
 * Example below:
 */

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, e.TotalLeaves-e.LeavesTaken)
}
