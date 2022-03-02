package main

import (
	"fmt"
	"sync"
)

/**
 * If one Goroutine already holds the lock and
 * if a new Goroutine is trying to acquire a lock,
 * the new Goroutine will be blocked until the mutex is unlocked.
**/

// 1. Race condition
/* var x = 0

func increament(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
} */

// 2. Solve race condition above using mutex
/* var x = 0

func increament(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
} */

// 3. Solving the race condition using channel
var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	// 1. Race condition
	/* var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increament(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x) */

	// 2. Solve race condition above using mutex
	/* var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increament(&w, &m) // passed as a pointer to reuse it so that each routine will not have its own mutex
	}
	w.Wait()
	fmt.Println("final value of x", x) */

	// 3. Solving the race condition using channel
	var w sync.WaitGroup
	ch := make(chan bool, 1) // must have size so that any other go routines trying to access x will wait until the 1 value is read from this cahnnel
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x is:", x)
}
