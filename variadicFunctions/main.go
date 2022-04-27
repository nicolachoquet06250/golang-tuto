package variadicFunctions

import "fmt"

/*func hello(...a int, b int) { // syntax error
	fmt.Println(a, b)
}*/
func hello(a int, b ...int) {
	fmt.Println(a, b)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findWithSlices(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func Main() {
	fmt.Println("\n-- Hello --")
	hello(1, 2)          //passing one argument "2" to b
	hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b

	fmt.Println("\n-- Find --")
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	nums := []int{89, 90, 95}
	// find(89, nums) // cannot use nums (type []int) as type int in argument to find
	find(89, nums...)

	fmt.Println("\n-- Find with slices --")
	findWithSlices(89, []int{89, 90, 95})
	findWithSlices(45, []int{56, 67, 45, 90, 109})
	findWithSlices(78, []int{38, 56, 98})
	findWithSlices(87, []int{})

	fmt.Println("\n-- Find with slices --")
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
