package loops

import "fmt"

func Main() {
	func() {
		fmt.Println("")
		for i := 1; i <= 10; i++ {
			fmt.Printf(" %d", i)
		}
	}()

	func() {
		fmt.Println("")
		for i := 1; i <= 10; i++ {
			if i > 5 {
				break //loop is terminated if i > 5
			}
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\nline after for loop")
	}()

	func() {
		fmt.Println("")
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				continue
			}
			fmt.Printf("%d ", i)
		}
	}()

	func() {
		fmt.Println("")
		n := 5
		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}()

	func() {
		fmt.Println("")
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
			}
		}
	}()

	func() {
		fmt.Println("")
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
				if i == j {
					break
				}
			}
		}
	}()

	func() {
		fmt.Println("")
	outer:
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
				if i == j {
					break outer
				}
			}
		}
	}()

	func() {
		fmt.Println("")
		i := 0
		for i <= 10 { // initialisation and post are omitted
			fmt.Printf("%d ", i)
			i += 2
		}
	}()

	func() {
		fmt.Println("")
		for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { // multiple initialisation and increment
			fmt.Printf("%d * %d = %d\n", no, i, no*i)
		}
	}()

	// infinite loop ( condition for finish program )
	func() {
		fmt.Println("")
		cmp := 0
		for {
			fmt.Println("Hello World")
			if cmp == 10 {
				break
			}
			cmp++
		}
	}()
}
