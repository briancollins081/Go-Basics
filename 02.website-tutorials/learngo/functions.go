package main

import "fmt"

// func functionname(parametername type) returntype {
// 	//function body
// }

// Single Return Value
func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func calculateBill2(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Multiple Return Values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = 2 * (length + width)
	return area, perimeter
}

//Named return types
func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length * width)
	return
}

func main() {
	// Single return
	price, no := 90, 16
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is ", totalPrice)

	// Multiple return
	area, perimeter := rectProps(10.5, 5.5)
	fmt.Println("The area and perimeter is ", area, perimeter)

	// Named return types
	area2, perimeter2 := rectProps(10.5, 5.5)
	fmt.Println("The area and perimeter is ", area2, perimeter2)

	// Blank identifier _
	_, perimeter3 := rectProps(230, 60)
	fmt.Printf("The perimeter is %2f\n", perimeter3)
}
