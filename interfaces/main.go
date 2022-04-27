package interfaces

import "fmt"

// VowelsFinder interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// FindVowels MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

// CalculateSalary salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// CalculateSalary salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("\nTotal Expense Per Month $%d", expense)
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

// CalculateSalary salary of freelancer
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

type Worker interface {
	Work()
}

type Person struct {
	name string
	age  int
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("\nInterface type %T value %v\n", w, w)
}

func assert(i interface{}) {
	/*s := i.(int) //get the underlying int value from i
	fmt.Println(s) //panic: interface conversion: interface {} is string, not int.*/
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func (p Person) Describe() {
	fmt.Printf("\n%s is %d years old", p.name, p.age)
}

type Describer interface {
	Describe()
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("\nState %s Country %s", a.state, a.country)
}

type EmployeeOperations interface {
	SalaryCalculator2
	LeaveCalculator
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

type SalaryCalculator2 interface {
	DisplaySalary()
}

func (e Employee) DisplaySalary() {
	fmt.Printf("\n%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func Main() {
	func() {
		name := MyString("Sam Anderson")
		var v VowelsFinder
		v = name // possible since MyString implements VowelsFinder
		fmt.Printf("Vowels are %c", v.FindVowels())
	}()

	func() {
		pemp1 := Permanent{
			empId:    1,
			basicpay: 5000,
			pf:       20,
		}
		pemp2 := Permanent{
			empId:    2,
			basicpay: 6000,
			pf:       30,
		}
		cemp1 := Contract{
			empId:    3,
			basicpay: 3000,
		}
		employees := []SalaryCalculator{pemp1, pemp2, cemp1}
		totalExpense(employees)
	}()

	func() {
		pemp1 := Permanent{
			empId:    1,
			basicpay: 5000,
			pf:       20,
		}
		pemp2 := Permanent{
			empId:    2,
			basicpay: 6000,
			pf:       30,
		}
		cemp1 := Contract{
			empId:    3,
			basicpay: 3000,
		}
		employees := []SalaryCalculator{pemp1, pemp2, cemp1}
		totalExpense(employees)
	}()

	func() {
		pemp1 := Permanent{
			empId:    1,
			basicpay: 5000,
			pf:       20,
		}
		pemp2 := Permanent{
			empId:    2,
			basicpay: 6000,
			pf:       30,
		}
		cemp1 := Contract{
			empId:    3,
			basicpay: 3000,
		}
		freelancer1 := Freelancer{
			empId:       4,
			ratePerHour: 70,
			totalHours:  120,
		}
		freelancer2 := Freelancer{
			empId:       5,
			ratePerHour: 100,
			totalHours:  100,
		}
		employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
		totalExpense(employees)
	}()

	func() {
		p := Person{
			name: "Naveen",
		}
		var w Worker = p
		describe(w)
		w.Work()
	}()

	func() {
		describe := func(i interface{}) {
			fmt.Printf("Type = %T, value = %v\n", i, i)
		}

		s := "Hello World"
		describe(s)
		i := 55
		describe(i)
		strt := struct {
			name string
		}{
			name: "Naveen R",
		}
		describe(strt)
	}()

	func() {
		var s interface{} = 56
		assert(s)
		var i interface{} = "Steven Paul"
		assert(i)
	}()

	func() {
		findType("Naveen")
		findType(77)
		findType(89.98)
	}()

	func() {
		findType("Naveen")
		p := Person{
			name: "Naveen R",
			age:  25,
		}
		findType(p)
	}()

	func() {
		var d1 Describer
		p1 := Person{"Sam", 25}
		d1 = p1
		d1.Describe()
		p2 := Person{"James", 32}
		d1 = &p2
		d1.Describe()

		var d2 Describer
		a := Address{"Washington", "USA"}

		/* compilation error if the following line is
		   uncommented
		   cannot use a (type Address) as type Describer
		   in assignment: Address does not implement
		   Describer (Describe method has pointer
		   receiver)
		*/
		//d2 = a

		d2 = &a //This works since Describer interface
		//is implemented by Address pointer in line 22
		d2.Describe()
	}()

	func() {
		e := Employee{
			firstName:   "Naveen",
			lastName:    "Ramanathan",
			basicPay:    5000,
			pf:          200,
			totalLeaves: 30,
			leavesTaken: 5,
		}
		var s SalaryCalculator2 = e
		s.DisplaySalary()
		var l LeaveCalculator = e
		fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
		var empOp EmployeeOperations = e
		empOp.DisplaySalary()
		fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
	}()

	func() {
		var d1 Describer
		if d1 == nil {
			fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
		}
	}()
}
