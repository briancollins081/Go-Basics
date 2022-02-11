// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func rot13(in rune) rune {
// 	if in >= 'A' && in <= 'Z' {
// 		return ((((in - 'A') + 13) % 26) + 'A')
// 	}

// 	if in > 'a' && in <= 'z' {
// 		return ((((in - 'a') + 13) % 26) + 'a')
// 	}

// 	return in
// }

// func main() {
// 	s := "This is a test 123 😀"

// 	s2 := strings.Map(rot13, s)
// 	fmt.Println(s2)

// 	s3 := strings.Map(rot13, s)
// 	fmt.Println(s3)
// }

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	s := "👋 🌏"
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(utf8.RuneCountInString(s))
// }

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Int31n(10)
	b := rand.Int31n(10)
	c := int(math.Max(float64(a), float64(b)))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, "is bigger")
}
