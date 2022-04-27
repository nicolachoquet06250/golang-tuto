package ifElse

import "fmt"

func Main() {
	num := 99
	if num <= 50 {
		fmt.Println(num, "is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println(num, "is between 51 and 100")
	} else {
		fmt.Println(num, "is greater than 100")
	}
	fmt.Println("The number", num, "is odd")

	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
