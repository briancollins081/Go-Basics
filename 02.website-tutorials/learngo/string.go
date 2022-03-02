package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Print("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printBytesUsingRunes(s string) {
	fmt.Print("Bytes: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printCharsUsingRune(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printBytesPosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func compareString(s1 string, s2 string) {
	if s1 == s2 {
		fmt.Printf("%s and %s are equal\n", s1, s2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", s1, s2)
}

/* func mutate(s string) string {
	s[0] = 'a' // any unicode character in single quotes becomes a rune
	return s
} */

func mutate(s []rune) string {
	s[0] = 'a' // any unicode character in single quotes becomes a rune
	return string(s)
}

func main() {
	// name := "Welcome to Kratos Island"
	/* println(name)
	// Accessing individual bytes
	printBytes(name)
	// characters
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n\n")
	name = "Señor®©"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Println("Using runes to print string")
	printCharsUsingRune(name)
	fmt.Println("Using for loop to print runes and their positions")
	printBytesPosition(name) */

	/* // Creating a string from an array of bytes
	bytesSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(bytesSlice)
	fmt.Println(str)
	// decimal equivqlents
	bytesSlice = []byte{67, 97, 102, 195, 169}
	str = string(bytesSlice)
	fmt.Println(str)
	//string from a slice of runes
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str) */

	// String Length
	/* word1 := "P®¥©.com"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Print("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2)) */

	// Compare two strings
	/* str1 := "Go"
	str2 := "Go"
	compareString(str1, str2)
	str1 = "Go"
	str2 = "golang"
	compareString(str1, str2) */

	// Concatenate strings
	/* str1 := "Go"
	str2 := "is awesome"
	result := str1 + " " + str2
	fmt.Println(result)
	fmt.Printf("%s %s", str1, str2) */

	// Mutate - you can not mutate a string once created
	/* h = "hello"
	fmt.Println(mutate(h)) */
	// workaround to mutate a string
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
