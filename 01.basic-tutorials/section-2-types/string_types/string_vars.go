package main

import "fmt"

func main() {
	/*
		STRINGS
	*/
	// var s string
	// s = "Pinto Gama Pilo"
	// s := "Hello\" Pinto Gama Pilo"
	// fmt.Println(s)

	// raw := `
	// 👏👏🌏
	// `
	// fmt.Println(raw + s)

	// s := "Hello Pinto Gama Pilo"
	// s2 := s[0:5]
	// s3 := s[7:12]
	// fmt.Println(s, len(s), s2, len(s2), s3, len(s3))

	// s := "👋 🌏"
	// s2 := s[:1]
	// s3 := s[2:]
	// fmt.Println(s, len(s), s2, len(s2), s3, len(s3))

	/*
		RUNES
	*/
	// s := "Chill Bro, still begining this "
	// var r rune = 127757
	// s = s + string(r)
	// fmt.Println(s)

	// s := "Chill Bro, still begining this "
	// r := '😘'
	// s = s + string(r)
	// fmt.Println(s)

	/*
		ARRAYS
	*/
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 8
	fmt.Println(vals, vals[0], vals[1], vals[2])
}
