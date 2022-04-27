package arraysSlices

import "fmt"

func Main() {
	fmt.Println("\n-- Arrays --")

	func() {
		var a [3]int //int array with length 3
		fmt.Println(a)
	}()
	func() {
		var a [3]int //int array with length 3
		a[0] = 12    // array index starts at 0
		a[1] = 78
		a[2] = 50
		fmt.Println(a)
	}()
	func() {
		a := [3]int{12, 78, 50} // short hand declaration to create array
		fmt.Println(a)
	}()
	func() {
		a := [3]int{12}
		fmt.Println(a)
	}()
	func() {
		a := [...]int{12, 78, 50} // ... makes the compiler determine the length
		fmt.Println(a)
	}()
	func() {
		//a := [3]int{5, 78, 8}
		//var b [5]int
		//b = a //not possible since [3]int and [5]int are distinct types
	}()
	func() {
		a := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a // a copy of a is assigned to b
		b[0] = "Singapore"
		fmt.Println("a is ", a)
		fmt.Println("b is ", b)
	}()
	func() {
		changeLocal := func(num [5]int) {
			num[0] = 55
			fmt.Println("inside function ", num)
		}

		num := [...]int{5, 6, 7, 8, 8}
		fmt.Println("before passing to function ", num)
		changeLocal(num) //num is passed by value
		fmt.Println("after passing to function ", num)
	}()
	func() {
		a := [...]float64{67.7, 89.8, 21, 78}
		fmt.Println("length of a is", len(a))
	}()
	func() {
		a := [...]float64{67.7, 89.8, 21, 78}
		for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
			fmt.Printf("%d th element of a is %.2f\n", i, a[i])
		}
	}()
	func() {
		a := [...]float64{67.7, 89.8, 21, 78}
		sum := float64(0)
		for i, v := range a { //range returns both the index and value
			fmt.Printf("%d the element of a is %.2f\n", i, v)
			sum += v
		}
		fmt.Println("\nsum of all elements of a", sum)
	}()
	func() {
		printarray := func(a [3][2]string) {
			for _, v1 := range a {
				for _, v2 := range v1 {
					fmt.Printf("%s ", v2)
				}
				fmt.Printf("\n")
			}
		}

		a := [3][2]string{
			{"lion", "tiger"},
			{"cat", "dog"},
			{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
		}

		printarray(a)

		var b [3][2]string
		b[0][0] = "apple"
		b[0][1] = "samsung"
		b[1][0] = "microsoft"
		b[1][1] = "google"
		b[2][0] = "AT&T"
		b[2][1] = "T-Mobile"

		fmt.Printf("\n")

		printarray(b)
	}()

	fmt.Println("\n-- Slices --")

	func() {
		a := [5]int{76, 77, 78, 79, 80}
		var b []int = a[1:4] //creates a slice from a[1] to a[3]
		fmt.Println(b)
	}()
	func() {
		c := []int{6, 7, 8} //creates and array and returns a slice reference
		fmt.Println(c)
	}()
	func() {
		darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		dslice := darr[2:5]
		fmt.Println("array before", darr)
		for i := range dslice {
			dslice[i]++
		}
		fmt.Println("array after", darr)
	}()
	func() {
		numa := [3]int{78, 79, 80}
		nums1 := numa[:] //creates a slice which contains all elements of the array
		nums2 := numa[:]
		fmt.Println("array before change 1", numa)
		nums1[0] = 100
		fmt.Println("array after modification to slice nums1", numa)
		nums2[1] = 101
		fmt.Println("array after modification to slice nums2", numa)
	}()
	func() {
		fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		fruitslice := fruitarray[1:3]
		fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
	}()
	func() {
		fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		fruitslice := fruitarray[1:3]
		fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
		fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
		fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
	}()
	func() {
		i := make([]int, 5, 5)
		fmt.Println(i)
	}()
	func() {
		cars := []string{"Ferrari", "Honda", "Ford"}
		fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
		cars = append(cars, "Toyota")
		fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	}()
	func() {
		var names []string //zero value of a slice is nil
		if names == nil {
			fmt.Println("slice is nil going to append")
			names = append(names, "John", "Sebastian", "Vinay")
			fmt.Println("names contents:", names)
		}
	}()
	func() {
		veggies := []string{"potatoes", "tomatoes", "brinjal"}
		fruits := []string{"oranges", "apples"}
		food := append(veggies, fruits...)
		fmt.Println("food:", food)
	}()
	func() {
		subtactOne := func(numbers []int) {
			for i := range numbers {
				numbers[i] -= 2
			}
		}

		nos := []int{8, 7, 6}
		fmt.Println("slice before function call", nos)
		subtactOne(nos)                               //function modifies the slice
		fmt.Println("slice after function call", nos) //modifications are visible outside
	}()
	func() {
		pls := [][]string{
			{"C", "C++"},
			{"JavaScript"},
			{"Go", "Rust"},
		}
		for _, v1 := range pls {
			for _, v2 := range v1 {
				fmt.Printf("%s ", v2)
			}
			fmt.Printf("\n")
		}
	}()
	func() {
		countries := func() []string {
			countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
			neededCountries := countries[:len(countries)-2]
			countriesCpy := make([]string, len(neededCountries))
			copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
			return countriesCpy
		}

		countriesNeeded := countries()
		fmt.Println(countriesNeeded)
	}()
}
