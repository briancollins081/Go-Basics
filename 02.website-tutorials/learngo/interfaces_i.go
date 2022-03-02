package main

import "fmt"

/*
	1. Simple interface implememntation
*/

/* /* import "fmt"

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// My String implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam is a guru")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
}
*/

/*
	2. We will write a simple program that calculates the total expense for a company based on the individual salaries of the employees.
	For brevity, we have assumed that all expenses are in USD.
*/

/* type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

// Frelancer addition without need to change totalExpenses
type Frelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

//salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicPay
}

//salary of freelancer
func (f Frelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}


// total expense is calculated by iterating through the SalaryCalculator slice and summing
// the salaries of the individual employees

func totalExpenses(s []SalaryCalculator) {
	expenses := 0
	for _, v := range s {
		expenses = expenses + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expenses)
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       50,
	}
	cemp1 := Permanent{
		empId:    3,
		basicPay: 4500,
	}
	frelancer1 := Frelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	frelancer2 := Frelancer{
		empId:       45,
		ratePerHour: 100,
		totalHours:  100,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1, frelancer1, frelancer2}

	totalExpenses(employees)
} */

/*
	3.Internal Interface representation
*/

/*
 type Worker interface {
	Work()
}
type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, "is working!")
}

func describe(w Worker) {
	fmt.Printf("interface type %T value %v\n", w, w)
}

func main() {
	p := Person{
		name: "Naveen",
	}
	var w Worker = p
	describe(w)
	w.Work()
} */

/*
	4.Empty interface
	  Since the empty interface has zero methods, all types implement the empty interface.
*/

/* func describe(i interface{}) {
	fmt.Printf("Type = %T, value= %v\n", i, i)
}

func main() {
	s := "Hello Andere"
	describe(s)
	i := 555
	describe(i)
	strt := struct {
		name string
	}{
		name: "Andere Brian",
	}
	describe(strt)
} */

/*
	5. Type Assertion
	   Used to extract the underlying value of the interface
*/
// func assert(i interface{}) {
// 	s := i.(int) // get the underlying int value from i
// 	fmt.Println(s)
// }

/* func assert(i interface{}) {
	j, okay := i.(int) // get the underlying int value from i
	fmt.Println(j, okay)
}
func main() {
	var s interface{} = 56
	assert(s)

	var i interface{} = "Banep DEO"
	assert(i)
} */

/*
	6. Type Switch
*/
/* func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string, and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int, and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	findType("Naveen")
	findType(9900)
	findType(990.9007)
} */

/*
	6. Type comparison to an interface
*/
type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type!\n")
	}
}

func main() {
	findType("Andere")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType(p)
}
