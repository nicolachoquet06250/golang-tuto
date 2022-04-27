package switchCase

import (
	"fmt"
	"math/rand"
)

func Main() {
	func() {
		finger := 4
		fmt.Printf("Finger %d is ", finger)
		switch finger {
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
		default: //default case
			fmt.Println("incorrect finger number")
		}
	}()

	func() {
		letter := "i"
		fmt.Printf("Letter %s is a ", letter)
		switch letter {
		case "a", "e", "i", "o", "u": //multiple expressions in case
			fmt.Println("vowel")
		default:
			fmt.Println("not a vowel")
		}
	}()

	func() {
		num := 75
		switch { // expression is omitted
		case num >= 0 && num <= 50:
			fmt.Printf("%d is greater than 0 and less than 50", num)
		case num >= 51 && num <= 100:
			fmt.Printf("%d is greater than 51 and less than 100", num)
		case num >= 101:
			fmt.Printf("%d is greater than 100", num)
		}
	}()

	func() {
		/*switch num := number(); { //num is not a constant
		case num < 50:
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough
		case num < 100:
			fmt.Printf("%d is lesser than 100\n", num)
			fallthrough
		case num < 200:
			fmt.Printf("%d is lesser than 200", num)
		}*/
	}()

	func() {
		switch num := 25; {
		case num < 50:
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough
		case num > 100:
			fmt.Printf("%d is greater than 100\n", num)
		}
	}()

	func() {
		switch num := -5; {
		case num < 50:
			if num < 0 {
				break
			}
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough
		case num < 100:
			fmt.Printf("%d is lesser than 100\n", num)
			fallthrough
		case num < 200:
			fmt.Printf("%d is lesser than 200", num)
		}
	}()

	func() {
	randloop:
		for {
			switch i := rand.Intn(100); {
			case i%2 == 0:
				fmt.Printf("Generated even number %d", i)
				break randloop
			}
		}
	}()
}
