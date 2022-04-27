package selects

import (
	"fmt"
	"time"
)

func Main() {
	func() {
		server1 := func(ch chan string) {
			time.Sleep(6 * time.Second)
			ch <- "from server1"
		}
		server2 := func(ch chan string) {
			time.Sleep(3 * time.Second)
			ch <- "from server2"
		}

		output1 := make(chan string)
		output2 := make(chan string)
		go server1(output1)
		go server2(output2)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		}
	}()

	func() {
		process := func(ch chan string) {
			time.Sleep(10500 * time.Millisecond)
			ch <- "process successful"
		}

		ch := make(chan string)
		go process(ch)
		for {
			time.Sleep(1000 * time.Millisecond)
			select {
			case v := <-ch:
				fmt.Println("received value: ", v)
				return
			default:
				fmt.Println("no value received")
			}
		}
	}()

	func() {
		ch := make(chan string)
		select {
		case <-ch:
		default:
			fmt.Println("default case executed")
		}
	}()

	func() {
		var ch chan string
		select {
		case v := <-ch:
			fmt.Println("received value", v)
		default:
			fmt.Println("default case executed")
		}
	}()

	func() {
		server1 := func(ch chan string) {
			ch <- "from server1"
		}
		server2 := func(ch chan string) {
			ch <- "from server2"
		}

		output1 := make(chan string)
		output2 := make(chan string)
		go server1(output1)
		go server2(output2)
		time.Sleep(1 * time.Second)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		}
	}()
}
