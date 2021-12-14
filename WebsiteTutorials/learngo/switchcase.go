package main

import (
	"fmt"
	"math/rand"
)

func main() {
	finger := 40
	fmt.Printf("Finger %d is ", finger)

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Ring")
	default:
		fmt.Println("incorrect finger number")
	}

	// Multiple expressions in case
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not vowel")
	}

	// Expressionless switch
	num := 75
	switch {
	case num > 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 51\n", num)
	case num > 51 && num <= 100:
		fmt.Printf("%d is greater than 50 and less than 101\n", num)
	case num >= 101:
		fmt.Printf("%d is greater than 101\n", num)
	}

	// Fallthrough case
	switch num := number(); {
	case num < 50:
		fmt.Printf("\n%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("\n%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("\n%d is lesser than 200\n", num)
	}
	fmt.Println("Test 2: Fall through behaviour")
	switch num = 25; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}

	// Breaking from a switch

	switch num = -5; {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("\n%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	// Breaking the outer loop
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("\nGenerated even number %d", i)
			break randloop
		}
	}
}

func number() int {
	num := 15 * 6
	return num
}
