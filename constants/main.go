package constants

import (
	"fmt"
)

func Main() {
	func() {
		const a = 50
		fmt.Println(a)
	}()

	func() {
		const (
			name    = "John"
			age     = 50
			country = "Canada"
		)
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(country)
	}()

	// error
	func() {
		const a = 55 //allowed
		//a = 89 //reassignment not allowed
	}()

	func() {
		//var a = math.Sqrt(4)   //allowed
		//const b = math.Sqrt(4) //not allowed
	}()

	func() {
		const n = "Sam"
		var name = n
		fmt.Printf("type %T value %v", name, name)
	}()

	func() {
		const a = 5
		var intVar int = a
		var int32Var int32 = a
		var float64Var float64 = a
		var complex64Var complex64 = a
		fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	}()

	func() {
		var i = 5
		var f = 5.6
		var c = 5 + 6i
		fmt.Printf("i's type is %T, f's type is %T, c's type is %T", i, f, c)
	}()

	func() {
		var a = 5.9 / 8
		fmt.Printf("a's type is %T and value is %v", a, a)
	}()
}
