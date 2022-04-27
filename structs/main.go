package structs

import "fmt"

func Main() {
	func() {
		type Employee struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}

		//creating struct specifying field names
		emp1 := Employee{
			firstName: "Sam",
			age:       25,
			salary:    500,
			lastName:  "Anderson",
		}

		//creating struct without specifying field names
		emp2 := Employee{"Thomas", "Paul", 29, 800}

		fmt.Println("Employee 1", emp1)
		fmt.Println("Employee 2", emp2)
	}()

	func() {
		emp3 := struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}{
			firstName: "Andreah",
			lastName:  "Nikola",
			age:       31,
			salary:    5000,
		}

		fmt.Println("Employee 3", emp3)
	}()

	func() {
		type Employee struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}

		emp6 := Employee{
			firstName: "Sam",
			lastName:  "Anderson",
			age:       55,
			salary:    6000,
		}
		fmt.Println("First Name:", emp6.firstName)
		fmt.Println("Last Name:", emp6.lastName)
		fmt.Println("Age:", emp6.age)
		fmt.Printf("Salary: $%d\n", emp6.salary)
		emp6.salary = 6500
		fmt.Printf("New Salary: $%d", emp6.salary)
	}()

	func() {
		type Employee struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}

		var emp4 Employee //zero valued struct
		fmt.Println("First Name:", emp4.firstName)
		fmt.Println("Last Name:", emp4.lastName)
		fmt.Println("Age:", emp4.age)
		fmt.Println("Salary:", emp4.salary)
	}()

	func() {
		type Employee struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}

		emp5 := Employee{
			firstName: "John",
			lastName:  "Paul",
		}
		fmt.Println("First Name:", emp5.firstName)
		fmt.Println("Last Name:", emp5.lastName)
		fmt.Println("Age:", emp5.age)
		fmt.Println("Salary:", emp5.salary)
	}()

	func() {
		type Employee struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}

		emp8 := &Employee{
			firstName: "Sam",
			lastName:  "Anderson",
			age:       55,
			salary:    6000,
		}
		fmt.Println("First Name:", (*emp8).firstName)
		fmt.Println("Age:", (*emp8).age)
	}()

	func() {
		type Employee struct {
			firstName string
			lastName  string
			age       int
			salary    int
		}

		emp8 := &Employee{
			firstName: "Sam",
			lastName:  "Anderson",
			age:       55,
			salary:    6000,
		}
		fmt.Println("First Name:", emp8.firstName)
		fmt.Println("Age:", emp8.age)
	}()

	func() {
		type Person struct {
			string
			int
		}

		p1 := Person{
			string: "naveen",
			int:    50,
		}
		fmt.Println(p1.string)
		fmt.Println(p1.int)
	}()

	func() {
		type Address struct {
			city  string
			state string
		}

		type Person struct {
			name    string
			age     int
			address Address
		}

		p := Person{
			name: "Naveen",
			age:  50,
			address: Address{
				city:  "Chicago",
				state: "Illinois",
			},
		}

		fmt.Println("Name:", p.name)
		fmt.Println("Age:", p.age)
		fmt.Println("City:", p.address.city)
		fmt.Println("State:", p.address.state)
	}()

	func() {
		type Address struct {
			city  string
			state string
		}
		type Person struct {
			name string
			age  int
			Address
		}

		p := Person{
			name: "Naveen",
			age:  50,
			Address: Address{
				city:  "Chicago",
				state: "Illinois",
			},
		}

		fmt.Println("Name:", p.name)
		fmt.Println("Age:", p.age)
		fmt.Println("City:", p.city)   //city is promoted field
		fmt.Println("State:", p.state) //state is promoted field
	}()

	func() {
		type name struct {
			firstName string
			lastName  string
		}

		name1 := name{
			firstName: "Steve",
			lastName:  "Jobs",
		}
		name2 := name{
			firstName: "Steve",
			lastName:  "Jobs",
		}
		if name1 == name2 {
			fmt.Println("name1 and name2 are equal")
		} else {
			fmt.Println("name1 and name2 are not equal")
		}

		name3 := name{
			firstName: "Steve",
			lastName:  "Jobs",
		}
		name4 := name{
			firstName: "Steve",
		}

		if name3 == name4 {
			fmt.Println("name3 and name4 are equal")
		} else {
			fmt.Println("name3 and name4 are not equal")
		}
	}()

	func() {
		/*type image struct {
			data map[int]int
		}

		image1 := image{
			data: map[int]int{
				0: 155,
			}}
		image2 := image{
			data: map[int]int{
				0: 155,
			}}
		if image1 == image2 {
		// invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
			fmt.Println("image1 and image2 are equal")
		}*/
	}()
}
