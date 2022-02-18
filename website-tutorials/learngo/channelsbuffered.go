package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 2. Example 2 - buffred channels
/* func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "to ch")
	}
	close(ch)
} */

// 5. WorkerPools -- WaitGroups
/* func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
} */

// 6. Worker pool implementation - use of buffered channels
type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(time.Second * 2)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
func main() {
	// 1. Example - buffered channels

	/* ch := make(chan string, 2) //Since the channel has a capacity of 2, it is possible to write 2 strings into the channel without being blocked.
	ch <- "Brian"
	ch <- "Collins"
	fmt.Println(<-ch)
	fmt.Println(<-ch) */

	// 2. Example 2
	/* ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	} */

	// 3. Closing buffered channels
	/* ch := make(chan int, 5)
	ch <- 50
	ch <- 68
	close(ch)
	n, open := <-ch
	fmt.Printf("Received %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received %d, open: %t\n", n, open)

	for n := range ch {
		fmt.Printf("Received %d\n", n)
	} */

	// 4. Length vs Capacity
	/* ch := make(chan string, 6)
	ch <- "halloween"
	ch <- "Id ul fitri"
	ch <- "Christmas"
	fmt.Println("Capacity is", cap(ch))
	fmt.Println("Length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch)) */

	// 5. WorkerPools -- WaitGroups
	/* no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finsished executing") */

	// 6. Worker pool implementation
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 1000
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
