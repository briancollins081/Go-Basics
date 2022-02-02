package main

import (
	"fmt"
	"structs/computers"
)

func main() {
	spec := computers.Spec{
		Maker: "Apple",
		Price: 5600,
		// model: "Chuck", //error
	}

	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
