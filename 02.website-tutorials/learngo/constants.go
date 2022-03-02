package main

import (
	"fmt"
)

func main() {
	const a = 50
	fmt.Println(a)
	// Group of constants
	const (
		name    = "John"
		age     = 78000
		country = "Kenya"
	)
	fmt.Println(name, age, country)

	// Types Constants
	// 1. Strings
	const namensis string = "Hello World"
	var justIntime = "Seim"
	type mulcherString string
	var customerName mulcherString = "Poller Bear"
	customerName = mulcherString(justIntime)
	fmt.Println(customerName)

	// 2.Boolean
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst
	defaultBool = bool(customBool)
	fmt.Println(defaultBool)

	// 3.Numbers
	const avar = 600000000
	var intVar int = avar
	var int32Var int32 = avar
	var float32Var float32 = avar
	var float64Var float64 = avar
	var complex64Var complex64 = avar
	var complex128Var complex128 = avar
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat32Var", float32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var, "\ncomplex128Var", complex128Var)
}
