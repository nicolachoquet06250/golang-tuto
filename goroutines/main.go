package goroutines

import (
	"fmt"
	"time"
)

func Main() {
	func() {
		hello := func() {
			fmt.Println("Hello world goroutine")
		}

		go hello()
		time.Sleep(1 * time.Second)
		fmt.Println("main function")
	}()

	func() {
		numbers := func() {
			for i := 1; i <= 5; i++ {
				time.Sleep(250 * time.Millisecond)
				fmt.Printf("%d ", i)
			}
		}
		alphabets := func() {
			for i := 'a'; i <= 'e'; i++ {
				time.Sleep(400 * time.Millisecond)
				fmt.Printf("%c ", i)
			}
		}

		go numbers()
		go alphabets()
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("main terminated")
	}()
}
