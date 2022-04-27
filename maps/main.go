package maps

import "fmt"

func Main() {
	func() {
		employeeSalary := make(map[string]int)
		fmt.Println(employeeSalary)
	}()

	func() {
		employeeSalary := make(map[string]int)
		employeeSalary["steve"] = 12000
		employeeSalary["jamie"] = 15000
		employeeSalary["mike"] = 9000
		fmt.Println("employeeSalary map contents:", employeeSalary)
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		employeeSalary["mike"] = 9000
		fmt.Println("employeeSalary map contents:", employeeSalary)
	}()

	func() {
		/*
			var employeeSalary map[string]int
			employeeSalary["steve"] = 12000 // panic: assignment to entry in nil map
		*/
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		fmt.Println("Salary of joe is", employeeSalary["joe"])
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		newEmp := "joe"
		value, ok := employeeSalary[newEmp]
		if ok == true {
			fmt.Println("Salary of", newEmp, "is", value)
			return
		}
		fmt.Println(newEmp, "not found")
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("Contents of the map")
		for key, value := range employeeSalary {
			fmt.Printf("employeeSalary[%s] = %d\n", key, value)
		}
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("map before deletion", employeeSalary)
		delete(employeeSalary, "steve")
		fmt.Println("map after deletion", employeeSalary)
	}()

	func() {
		type employee struct {
			salary  int
			country string
		}

		emp1 := employee{
			salary:  12000,
			country: "USA",
		}
		emp2 := employee{
			salary:  14000,
			country: "Canada",
		}
		emp3 := employee{
			salary:  13000,
			country: "India",
		}
		employeeInfo := map[string]employee{
			"Steve": emp1,
			"Jamie": emp2,
			"Mike":  emp3,
		}

		for name, info := range employeeInfo {
			fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
		}
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		fmt.Println("length is", len(employeeSalary))
	}()

	func() {
		employeeSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("Original employee salary", employeeSalary)
		modified := employeeSalary
		modified["mike"] = 18000
		fmt.Println("Employee salary changed", employeeSalary)
	}()
}
