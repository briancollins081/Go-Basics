package main

import "fmt"

/* type Employee struct {
	name     string
	age      int
	salary   int
	currency string
}

type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius float64
} */

// This method has Employee as the receiver type
/* func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
} */

// displaySalary() method converted to function with Employee as parameter
/* func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
} */

// Using same method name on different objects
/* func (r Rectangle) Area() int {
	return r.length * r.width
}
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
} */

// Pointer vs Value Recievers
/* func (e Employee) changeName(newName string) {
	e.name = newName
}
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
} */

// Methods ofanonymous structs
/* type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
} */

// Value receivers in methods(pointers & values) vs Value arguments in functions (only values)
/* type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area method result: %d\n", (r.length * r.width))
} */

// Pointer receivers in methods vs Pointer arguments in functions
/* type rectangle struct {
	length int
	width  int
}

func perimeter(r *rectangle) {
	fmt.Printf("perimeter function output: %d\n", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Printf("perimeter function output: %d\n", 2*(r.length+r.width))
} */

// Methods with non-struct receivers
type myInt int

func (a myInt) add(b myInt) myInt {
	return (a + b)
}
func main() {
	/* emp1 := Employee{
		name:     "Newett Adolf",
		salary:   5600,
		currency: "$",
	}
	// emp1.displaySalary()
	displaySalary(emp1) */

	// Using same method name on different objects
	/* r := Rectangle{length: 10, width: 5}
	fmt.Printf("The area of the triangle is %d\n", r.Area())
	c := Circle{radius: 12}
	fmt.Printf("The area of the triangle is %f\n", c.Area()) */

	// Pointer vs Value Recievers
	/* emp1 := Employee{
		name:     "Brian Andere",
		salary:   5600,
		currency: "$",
		age:      40,
	}
	fmt.Printf("Employee name before change: %s\n", emp1.name)
	emp1.changeName("Robinson Okoth")
	fmt.Printf("Employee name after 1st change: %s\n", emp1.name)
	emp1.changeName("Gibert Titusu")
	fmt.Printf("Employee name after 2nd change: %s\n", emp1.name)

	fmt.Printf("Employee age before change: %d\n", emp1.age)
	(&emp1).changeAge(34)
	fmt.Printf("Employee age after 1st change: %d\n", emp1.age)
	emp1.changeAge(60)
	fmt.Printf("Employee age after 2nd change: %d\n", emp1.age) */

	// Methods of anonymous struct fields
	/* p := person{
		firstName: "Elohim",
		lastName:  "Munchin",
		address: address{
			city:  "Nairobi",
			state: "City of Light",
		},
	}
	p.address.fullAddress() */

	// Value receivers in methods(pointers & values) vs Value arguments in functions (only values)
	/* r := rectangle{
		length: 56,
		width:  10,
	}
	area(r)
	r.area()
	p := &r
	// area(p) // error
	p.area() */

	// Pointer receivers in methods vs Pointer arguments in functions
	/* r := rectangle{
		length: 56,
		width:  10,
	}
	p := &r
	perimeter(p)
	p.perimeter()
	// perimeter(r) //error
	r.perimeter() */

	// Methods with non-struct receivers
	// To use literal types as receivers they must exist in the same package hence -> type myInt int
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
