package methods

import (
	"fmt"
	"math"
)

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("\nFull address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

type Employee struct {
	name     string
	salary   int
	currency string
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("\nSalary of %s is %s%d", e.name, e.currency, e.salary)
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (e Employee) changeName(name string) {
	e.name = name
}

func (e Employee) changeSalary(salary int) {
	e.salary = salary
}

func (e *Employee) changeSalaryWithPointer(salary int) {
	e.salary = salary
}

func (e Employee) changeCurrency(currency string) {
	e.currency = currency
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func Main() {
	func() {
		displaySalary := func(e Employee) {
			fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
		}

		emp1 := Employee{
			name:     "Sam Adolf",
			salary:   5000,
			currency: "$",
		}
		emp1.displaySalary() //Calling displaySalary() method of Employee type
		displaySalary(emp1)
	}()

	func() {
		r := Rectangle{
			length: 10,
			width:  5,
		}
		fmt.Printf("Area of rectangle %d\n", r.Area())
		c := Circle{
			radius: 12,
		}
		fmt.Printf("Area of circle %f", c.Area())
	}()

	func() {
		e := Employee{
			name:     "Sam Adolf",
			salary:   5_000,
			currency: "$",
		}

		fmt.Printf("\nEmployee name before change: %s", e.name)
		e.changeName("Michael Andrew")
		fmt.Printf("\nEmployee name after change: %s", e.name)

		fmt.Printf("\n\nEmployee salary before change: %d", e.salary)
		(&e).changeSalary(6_000)
		fmt.Printf("\nEmployee salary after change: %d", e.salary)
	}()

	func() {
		e := Employee{
			name:     "Sam Adolf",
			salary:   5_000,
			currency: "$",
		}

		fmt.Printf("\nEmployee name before change: %s", e.name)
		e.changeName("Michael Andrew")
		fmt.Printf("\nEmployee name after change: %s", e.name)

		fmt.Printf("\n\nEmployee salary before change: %d", e.salary)
		e.changeSalaryWithPointer(6_000)
		fmt.Printf("\nEmployee salary after change: %d", e.salary)
	}()

	func() {
		p := person{
			firstName: "Elon",
			lastName:  "Musk",
			address: address{
				city:  "Los Angeles",
				state: "California",
			},
		}

		p.fullAddress() //accessing fullAddress method of address struct
	}()

	func() {
		r := rectangle{
			length: 10,
			width:  5,
		}
		area(r)
		r.area()

		p := &r
		/*
		   compilation error, cannot use p (type *rectangle) as type rectangle
		   in argument to area
		*/
		//area(p)

		p.area() //calling value receiver with a pointer
	}()

	func() {
		r := rectangle{
			length: 10,
			width:  5,
		}
		p := &r //pointer to r
		perimeter(p)
		p.perimeter()

		/*
		   cannot use r (type rectangle) as type *rectangle in argument to perimeter
		*/
		//perimeter(r)

		r.perimeter() //calling pointer receiver with a value
	}()

	func() {
		num1 := myInt(5)
		num2 := myInt(10)
		sum := num1.add(num2)
		fmt.Println("Sum is", sum)
	}()
}
