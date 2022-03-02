package main

import (
	"fmt"
	"sync"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished() // called just before the largest function returns
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	fmt.Println("Largest number in", nums, "is", max)
}

// 2. Deferred methods
type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}

// 3. Arguments evaluation are done when defer is called
func printA(a int) {
	fmt.Println("value of a diferred function", a)
}

// 4. Stack of defers

// 5. Practical use of defer
type rect struct {
	length int
	width  int
}

/* func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero]\n", r)
		wg.Done()
		return
	}

	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero]\n", r)
		wg.Done()
		return
	}

	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
} */
// Rewrite the above with defer
func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero]\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero]\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	// 1.
	/* nums := []int{77, 34, 65, 90, 21, 56, 91, 89, 85}
	largest(nums) */

	// 2.
	/* p := person{
		firstName: "John",
		lastName:  "Collins",
	}
	defer p.fullName()
	fmt.Print("Welcome ") */

	// 3.
	/* a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a) */

	// 4. Stack of defers
	/* name := "Andere"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	} */

	// 5. Practical use of defer
	var wg sync.WaitGroup
	r1 := rect{91, 45}
	r2 := rect{60, -90}
	r3 := rect{50, 40}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished running")
}
