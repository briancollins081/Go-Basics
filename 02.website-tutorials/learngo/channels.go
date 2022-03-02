package main

import "fmt"

// 1. Example for channels
/* func hello(done chan bool) {
	fmt.Println("Hello go routine is going to sleep!")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello Go routine is awake, going to write to done")
	done <- true
} */

// 2. Example 2 - combining 2 channels
/* func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
} */

// 3.Unidirectional Channels
/* func sendData(sendch chan<- int) {
	sendch <- 10
} */

// 4. Closing channels and for range loops on channels
/* func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
} */

// 5. Using for range loop ex 3
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	/* var a chan int
	if a == nil {
		fmt.Println("channel a is nill, defining it in a few minutes")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	} */

	// 1.
	/* done := make(chan bool)
	fmt.Println("Main going to call go routine")
	go hello(done)
	<-done
	fmt.Println("Main recieved data") */

	// 2.
	/* number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes) */

	// 3.
	/* sendch := make(chan<- int)
	go sendData(sendch)
	// fmt.Println(<-sendch) // invalid operation: <-sendch (receive from send-only type chan<- int) */
	/* chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl) */

	// 4.
	/* ch := make(chan int)
	go producer(ch) */
	// using infinite for loop
	/* for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Recieved ", v, ok)
	} */
	// using for range to get value from a channel
	/* for v := range ch {
		fmt.Println("Recieved ", v)
	} */

	// 5.
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
