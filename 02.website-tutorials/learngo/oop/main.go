/**
 * Go does not provide classes but it does provide structs.
 * Methods can be added on structs.
 * This provides the behaviour of bundling the data and methods that
 * operate on the data together akin to a class.
 */
package main

import (
	"oop/employee"
)

func main() {
	// 1. Structs Instead of Classes
	/* e := employee.Employee{
		FirstName:   "Andere",
		LastName:    "Brian",
		TotalLeaves: 50,
		LeavesTaken: 35,
	}
	e.LeavesRemaining() */

	// 2. New() function instead of constructors
	/* var e employee.Employee  // returns a zero value
	e.LeavesRemaining() */

	// 3. using New() function
	e := employee.New("Brian", "Collins", 30, 25)
	e.LeavesRemaining()
}
