package pointers

import "fmt"

func Main() {
	func() {
		b := 255
		var a *int = &b
		fmt.Printf("Type of a is %T\n", a)
		fmt.Println("address of b is", a)
	}()

	func() {
		a := 25
		var b *int
		if b == nil {
			fmt.Println("b is", b)
			b = &a
			fmt.Println("b after initialization is", b)
		}
	}()

	func() {
		size := new(int)
		fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
		*size = 85
		fmt.Println("New size value is", *size)
	}()

	func() {
		b := 255
		a := &b
		fmt.Println("address of b is", a)
		fmt.Println("value of b is", *a)
	}()

	func() {
		b := 255
		a := &b
		fmt.Println("address of b is", a)
		fmt.Println("value of b is", *a)
		*a++
		fmt.Println("new value of b is", b)
	}()

	func() {
		change := func(val *int) {
			*val = 55
		}

		a := 58
		fmt.Println("value of a before function call is", a)
		b := &a
		change(b)
		fmt.Println("value of a after function call is", a)
	}()

	func() {
		hello := func() *int {
			i := 5
			return &i
		}

		d := hello()
		fmt.Println("Value of d", *d)
	}()

	func() {
		modify := func(arr *[3]int) {
			(*arr)[0] = 90
		}

		a := [3]int{89, 90, 91}
		modify(&a)
		fmt.Println(a)
	}()

	func() {
		modify := func(arr *[3]int) {
			arr[0] = 90
		}

		a := [3]int{89, 90, 91}
		modify(&a)
		fmt.Println(a)
	}()

	func() {
		modify := func(sls []int) {
			sls[0] = 90
		}

		a := [3]int{89, 90, 91}
		modify(a[:])
		fmt.Println(a)
	}()

	/*func() {
		b := [...]int{109, 110, 111}
		p := &b
		p++ // invalid operation: p++ (non-numeric type *[3]int)
	}()*/
}
