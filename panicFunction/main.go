package panicFunction

import (
	"fmt"
)

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		//debug.PrintStack()
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	/*defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	// divide(a, b, done) // recovered: runtime error: integer divide by zero
	<-done*/
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true
}

func Main() {
	func() {
		defer fmt.Println("deferred call in main")
		firstName := "Elon"
		fullName(&firstName, nil)
		fmt.Println("returned normally from main")
	}()

	func() {
		invalidSliceAccess()
		fmt.Println("normally returned from main")
	}()

	func() {
		sum(5, 0)
		fmt.Println("normally returned from main")
	}()
}
