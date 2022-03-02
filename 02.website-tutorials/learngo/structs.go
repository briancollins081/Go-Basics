package main

// Create a named stack
/* type Employee struct {
	firstName string
	lastName  string
	age       int
	salary	  int
} */
// or
/* type Employee struct {
	firstName, lastName string
	age, salary         int
} */

// Anonymous fields
/* type Person struct {
	string
	int
} */

// Nested structs
/* type Address struct {
	city  string
	state string
}
type Person struct {
	name    string
	age     int
	address Address
} */

// Promoted structs
/* type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	Address // anonymous field
} */

// Structs Equality
/* type name struct {
	firstName string
	lastName  string
} */

// Struct with uncompareable data
type image struct {
	data map[int]int
}

func main() {
	// Named Structs
	/* // create structby specifying field names
	emp1 := Employee{firstName: "Sam", age: 45, salary: 6000, lastName: "Collins"}
	// creating struct without specifying field names
	emp2 := Employee{"Thomas", "Paul", 29, 4500}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2) */

	// Anonymous Struct
	/* emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Poolenter",
		lastName:  "Nikoler",
		age:       31,
		salary:    32,
	}
	fmt.Println("Employee 3", emp3) */

	// Access struct elements
	/* emp6 := Employee{
		firstName: "Julius",
		lastName:  "Ceasar",
		age:       34,
		salary:    4000,
	}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
	emp6.salary = 6500
	fmt.Printf("New Salary: $%d", emp6.salary) */

	// Zero value of a struct
	/* var emp7 Employee //zero valued struct
	fmt.Println("First Name:", emp7.firstName)
	fmt.Println("Last Name:", emp7.lastName)
	fmt.Println("Age:", emp7.age)
	fmt.Printf("Salary: $%d\n", emp7.salary) */

	// Partial values
	/* emp8 := Employee{
		firstName: "Julius",
		lastName:  "Hagutt",
	}
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Last Name:", emp8.lastName)
	fmt.Println("Age:", emp8.age)
	fmt.Printf("Salary: $%d\n", emp8.salary) */

	// Pointers to a struct
	/* emp9 := &Employee{
		firstName: "Charles",
		lastName:  "Munich",
		age:       56,
		salary:    3000,
	}
	fmt.Println("Last Name:", (*emp9).lastName)
	fmt.Println("Salary: $", (*emp9).salary)
	// or
	fmt.Println("First Name:", emp9.firstName)
	fmt.Println("Age:", emp9.age) */

	// Anonymous fields
	/* p1 := Person{
		string: "navreel",
		int:    90007,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int) */

	// Nested Structs
	/* p := Person{
		name: "Andere Brian",
		age:  50,
		address: Address{
			city:  "Nairobi",
			state: "Nairobi",
		},
	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.address.city)
	fmt.Println("State: ", p.address.state) */

	// Promoted field
	/* p := Person{
		name: "Julius Ceasar",
		age:  50,
		Address: Address{
			city:  "Neopolis",
			state: "Neo",
		},
	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.city)   //promoted field
	fmt.Println("State: ", p.state) //promoted field */

	// Structs Equality
	/* name1 := name{
		firstName: "Steven",
		lastName:  "Joys",
	}
	name2 := name{
		firstName: "Steven",
		lastName:  "Joys",
	}

	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steven",
		lastName:  "Joys",
	}
	name4 := name{
		firstName: "Steven",
	}
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name1 and name4 are not equal")
	} */

	// Struct Variables are notcompareable if they contain fields that are not comparable
	/* image1 := image{
		data: map[int]int{
			0: 550,
		},
	}

	image2 := image{
		data: map[int]int{
			0: 550,
		},
	}

	if image1 == image2 { //fails
		fmt.Println("image1 is equal to image2")
	} */

}
