package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Main() {
	func() {
		var a chan int
		if a == nil {
			fmt.Println("channel a is nil, going to define it")
			a = make(chan int)
			fmt.Printf("Type of a is %T", a)
		}
	}()

	func() {
		hello := func(done chan bool) {
			fmt.Println("\nHello world goroutine")
			done <- true
		}

		done := make(chan bool)
		go hello(done)
		<-done
		fmt.Println("main function")
	}()

	func() {
		hello := func(done chan bool) {
			fmt.Println("hello go routine is going to sleep")
			time.Sleep(4 * time.Second)
			fmt.Println("hello go routine awake and going to write to done")
			done <- true
		}

		done := make(chan bool)
		fmt.Println("Main going to call hello go goroutine")
		go hello(done)
		<-done
		fmt.Println("Main received data")
	}()

	func() {
		calcSquares := func(number int, squareop chan int) {
			sum := 0
			for number != 0 {
				digit := number % 10
				sum += digit * digit
				number /= 10
			}
			squareop <- sum
		}

		calcCubes := func(number int, cubeop chan int) {
			sum := 0
			for number != 0 {
				digit := number % 10
				sum += digit * digit * digit
				number /= 10
			}
			cubeop <- sum
		}

		number := 589
		sqrch := make(chan int)
		cubech := make(chan int)
		go calcSquares(number, sqrch)
		go calcCubes(number, cubech)
		squares, cubes := <-sqrch, <-cubech
		fmt.Println("Final output", squares+cubes)
	}()

	func() {
		sendData := func(sendch chan<- int) {
			sendch <- 10
		}

		sendch := make(chan int)
		go sendData(sendch)
		// fmt.Println(<-sendch) // invalid operation: <-sendch (receive from send-only type chan<- int)
		fmt.Println(<-sendch)
	}()

	func() {
		producer := func(chnl chan int) {
			for i := 0; i < 10; i++ {
				chnl <- i
			}
			close(chnl)
		}

		ch := make(chan int)
		go producer(ch)
		for {
			v, ok := <-ch
			if ok == false {
				break
			}
			fmt.Println("Received ", v, ok)
		}
	}()

	func() {
		producer := func(chnl chan int) {
			for i := 0; i < 10; i++ {
				chnl <- i
			}
			close(chnl)
		}

		ch := make(chan int)
		go producer(ch)
		for v := range ch {
			fmt.Println("Received ", v)
		}
	}()

	func() {
		digits := func(number int, dchnl chan int) {
			for number != 0 {
				digit := number % 10
				dchnl <- digit
				number /= 10
			}
			close(dchnl)
		}

		calcSquares := func(number int, squareop chan int) {
			sum := 0
			dch := make(chan int)
			go digits(number, dch)
			for digit := range dch {
				sum += digit * digit
			}
			squareop <- sum
		}

		calcCubes := func(number int, cubeop chan int) {
			sum := 0
			dch := make(chan int)
			go digits(number, dch)
			for digit := range dch {
				sum += digit * digit * digit
			}
			cubeop <- sum
		}

		number := 589
		sqrch := make(chan int)
		cubech := make(chan int)
		go calcSquares(number, sqrch)
		go calcCubes(number, cubech)
		squares, cubes := <-sqrch, <-cubech
		fmt.Println("Final output", squares+cubes)
	}()

	fmt.Println("\n- Buffered Channels and Worker Pools -")

	func() {
		ch := make(chan string, 2)
		ch <- "naveen"
		ch <- "paul"
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	func() {
		write := func(ch chan int) {
			for i := 0; i < 5; i++ {
				ch <- i
				fmt.Println("successfully wrote", i, "to ch")
			}
			close(ch)
		}

		ch := make(chan int, 2)
		go write(ch)
		time.Sleep(2 * time.Second)
		for v := range ch {
			fmt.Println("read value", v, "from ch")
			time.Sleep(2 * time.Second)
		}
	}()

	func() {
		ch := make(chan int, 5)
		ch <- 5
		ch <- 6
		close(ch)
		n, open := <-ch
		fmt.Printf("Received: %d, open: %t\n", n, open)
		n, open = <-ch
		fmt.Printf("Received: %d, open: %t\n", n, open)
		n, open = <-ch
		fmt.Printf("Received: %d, open: %t\n", n, open)
	}()

	func() {
		ch := make(chan int, 5)
		ch <- 5
		ch <- 6
		close(ch)
		for n := range ch {
			fmt.Println("Received:", n)
		}
	}()

	func() {
		ch := make(chan string, 3)
		ch <- "naveen"
		ch <- "paul"
		fmt.Println("capacity is", cap(ch))
		fmt.Println("length is", len(ch))
		fmt.Println("read value", <-ch)
		fmt.Println("new length is", len(ch))
	}()

	func() {
		process := func(i int, wg *sync.WaitGroup) {
			fmt.Println("started Goroutine ", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d ended\n", i)
			wg.Done()
		}

		no := 3
		var wg sync.WaitGroup
		for i := 0; i < no; i++ {
			wg.Add(1)
			go process(i, &wg)
		}
		wg.Wait()
		fmt.Println("All go routines finished executing")
	}()

	func() {
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

		digits := func(number int) int {
			sum := 0
			no := number
			for no != 0 {
				digit := no % 10
				sum += digit
				no /= 10
			}
			time.Sleep(2 * time.Second)
			return sum
		}
		worker := func(wg *sync.WaitGroup) {
			for job := range jobs {
				output := Result{job, digits(job.randomno)}
				results <- output
			}
			wg.Done()
		}
		createWorkerPool := func(noOfWorkers int) {
			var wg sync.WaitGroup
			for i := 0; i < noOfWorkers; i++ {
				wg.Add(1)
				go worker(&wg)
			}
			wg.Wait()
			close(results)
		}
		allocate := func(noOfJobs int) {
			for i := 0; i < noOfJobs; i++ {
				randomno := rand.Intn(999)
				job := Job{i, randomno}
				jobs <- job
			}
			close(jobs)
		}
		result := func(done chan bool) {
			for result := range results {
				fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
			}
			done <- true
		}

		startTime := time.Now()
		noOfJobs := 100
		go allocate(noOfJobs)
		done := make(chan bool)
		go result(done)
		noOfWorkers := 10
		createWorkerPool(noOfWorkers)
		<-done
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		fmt.Println("total time taken ", diff.Seconds(), "seconds")
	}()
}
