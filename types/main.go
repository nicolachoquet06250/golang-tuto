package types

import (
	"fmt"
	"unsafe"
)

func boolean() {
	fmt.Println("\n-- Booleans --")
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d:", d)
}

func signedInteger() {
	fmt.Println("\n-- Signed Integer --")

	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)

	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b
}

func unsignedInteger() {
	fmt.Println("\n-- Unsigned Integer --")

	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

func complexTypes() {
	fmt.Println("\n-- Complex Types --")

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

func stringType() {
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last

	fmt.Println("My name is", name)
}

func typeCasting() {
	func() {
		i := 55   //int
		j := 67.8 //float64
		// sum := i + j //int + float64 not allowed
		sum := i + int(j)
		fmt.Println(sum)
	}()

	func() {
		i := 10
		var j float64 = float64(i) //this statement will not work without explicit conversion
		fmt.Println("j", j)
	}()
}

func Main() {
	boolean()
	signedInteger()
	unsignedInteger()
	complexTypes()
	stringType()
	typeCasting()
}
