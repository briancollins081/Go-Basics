package main

import "fmt"

func main() {
	/* employeeSalary := make(map[string]int)
	employeeSalary["mark"] = 12000
	employeeSalary["jamie"] = 12050
	employeeSalary["mike"] = 14000 */

	// Initialize a mapduring decalration
	/* 	employeeSalary := map[string]int{"mark": 12000, "jamie": 12050}
	   	employeeSalary["mike"] = 14000
	   	fmt.Println("employeeSalary map contents:", employeeSalary) */

	//zero value
	/* var employeeSalary map[string]int
	employeeSalary["steve"] = 899000 */

	// Reading a value
	/* employeeSalary := map[string]int{"steve": 12000, "jamie": 15000, "mike": 90000}
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
	employee = "ken"
	salary = employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary) */

	// Check if key exists
	/* newEmp := "piper"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "is not found!") */

	// Iterate over all elements
	/* fmt.Println("Content of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	} */

	// Deleting elements from a map
	/* fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deleteion", employeeSalary) */

	/* // Map of structs
	emp1 := employee{
		salary:  12000,
		country: "UK",
	}
	emp2 := employee{
		salary:  45000,
		country: "USA",
	}
	emp3 := employee{salary: 13000, country: "India"}
	emp4 := employee{salary: 16000, country: "Afghanistan"}
	employeeInfo := map[string]employee{
		"Steve":   emp1,
		"Jamie":   emp2,
		"Mike":    emp3,
		"Eduardo": emp4,
	}
	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary: %d Country: %s\n", name, info.salary, info.country)
	}
	// Length of map
	fmt.Println("Length is ", len(employeeInfo)) */

	// Maps are reference types
	employeeSalary := make(map[string]int)
	employeeSalary["mark"] = 12000
	employeeSalary["jamie"] = 12050
	employeeSalary["mike"] = 14000
	fmt.Println("Original employee salary:", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 5000000
	fmt.Println("Employee salary changed", employeeSalary)

	// Note: Map can only be compared to nil
}

type employee struct {
	salary  int
	country string
}
