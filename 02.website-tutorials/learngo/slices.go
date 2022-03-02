package main

import "fmt"

func main() {
	// Creating a slice
	//   1. from an array
	/* a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b) */

	//   2. Initialize
	/* c := []int{6, 7, 8} // creates an array and retruns a slice reference
	fmt.Println(c) */

	// Modifying a slice
	/* darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before ", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr) */

	// Changes affects both slices if they are using the same underlying array
	/* numa := [3]int{78, 79, 80}
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1 ", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1 ", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa) */

	// Length & Capacity
	/* fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)] // reslicing fruitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice)) */

	//  Creating a slice using make
	/* i := make([]int, 5, 5)
	fmt.Println(i) */

	// Appending to an array
	/* cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars: ", cars, " has old length ", len(cars))
	cars = append(cars, "Toyota", "Nissan")
	fmt.Println("cars: ", cars, " has new length ", len(cars)) */

	// Zero value of slice
	/* var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	} */

	// Appending two slices with ...
	/* veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food: ", food) */

	// Passing a slice to a function
	/* nos := []int{8, 7, 6}
	fmt.Println("slice before function call ", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos) */

	// Multidimensional slices
	/* pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	} */

	// Memory optimization
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy //Now countries array can be garbage collected since neededCountries is no longer referenced.
}
