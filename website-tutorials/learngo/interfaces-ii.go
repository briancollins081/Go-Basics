package main

import "fmt"

// 1.Value and Pointer Recivers
/* type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value reciever
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { // Using pointer reciever
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}*/

// 2.Implementing multiple interfaces

/* type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has a salary %d", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
} */

//3.Embedding interfaces
/* type SalaryCalculator interface {
	DisplaySalary()
}
type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has a salary of $%d", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
} */

// 4. Zero value
type Describer interface {
	Describe()
}

func main() {
	// 1.

	/* var d1 Describer
	p1 := Person{
		"Brian",
		30,
	}
	d1 = p1
	d1.Describe()
	p2 := Person{
		"Andere",
		25,
	}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"NewYork", "USA"}
	// d2 = a // ERROR: Describe method has pointer receiver
	d2 = &a // works since Describe method uses pointer reciever
	d2.Describe() */

	// 2.
	/* e := Employee{
		firstName:   "Benson",
		lastName:    "Black",
		basicPay:    9000,
		pf:          500,
		totalLeaves: 45,
		leavesTaken: 10,
	}

	var s SalaryCalculator = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft()) */

	// 3.
	/* e := Employee{
		firstName:   "Ruth",
		lastName:    "Jane",
		pf:          180,
		basicPay:    6500,
		totalLeaves: 45,
		leavesTaken: 21,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft()) */

	// 4.
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nill and has a type of %T value %v\n", d1, d1)
	}
}
