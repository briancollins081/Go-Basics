package main

// 1. Example of a Select
/* func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
} */

// 2. Default case
/* func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process succesful"
} */

// 4. Random selection
func server1(ch chan string) {
	ch <- "from server 1"
}

func server2(ch chan string) {
	ch <- "from server 2"
}

func server3(ch chan string) {
	ch <- "from server 3"
}

func main() {
	// 1.
	/* output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	} */

	// 2.
	/* ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("recieved value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	} */

	// 3. i Deadlock and default case
	// without the default block, the program will throw an error
	/* ch := make(chan string)
	select {
	case <-ch:
		fmt.Println("Recieved channel data!")
	default:
		fmt.Println("Will print, if ch does not recieve any data!")
	} */
	// 3. ii Using a nil channel
	/* var ch chan string
	select {
	case v := <-ch:
		fmt.Println("Received value:", v)
	default:
		fmt.Println("Default case executed")
	} */

	// 4. Random selection
	/* output1 := make(chan string)
	output2 := make(chan string)
	output3 := make(chan string)
	go server1(output1)
	go server2(output2)
	go server3(output3)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	case s3 := <-output3:
		fmt.Println(s3)
	} */

	// 5.Gotcha - Empty select - deadlock
	select {}
}
