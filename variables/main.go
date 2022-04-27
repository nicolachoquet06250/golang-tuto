package variables

import (
	"fmt"
	"math"
)

func firstStep() {
	var age int = 29 // variable declaration
	fmt.Println("My age is", age)

	age = 54
	fmt.Println("My new age is", age)
}

func secondStep() {
	var width, height int = 100, 50

	fmt.Println("width is", width, "height is", height)

	width = 100
	height = 50

	fmt.Println("new width is", width, "new height is", height)
}

func thirdStep() {
	var (
		name   = "naveen"
		age    = 29
		height int
	)

	fmt.Println("my name is", name, ", age is", age, "and height is", height)
}

func fourthStep() {
	count := 10

	fmt.Println("Count =", count)
}

func fifthStep() {
	name, age := "naveen", 29 //short hand declaration

	fmt.Println("my name is", name, "age is", age)
}

func sixthStep() {
	/*name, age := "naveen" //error

	fmt.Println("my name is", name, "age is", age)*/

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	/*a, b := 20, 30 //a and b declared
	fmt.Println("a is", a, "b is", b)
	a, b := 40, 50 //error, no new variables*/
}

func seventhStep() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)
}

func Main() {
	firstStep()
	secondStep()
	thirdStep()
	fourthStep()
	fifthStep()
	sixthStep()
	seventhStep()
}
