package structsClasses

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func NewEmployee(firstName string, lastName string, totalLeave int, leavesTaken int) Employee {
	return Employee{
		firstName,
		lastName,
		totalLeave,
		leavesTaken,
	}
}

func (e Employee) LeavesRemaining() {
	fmt.Printf(
		"%s %s has %d leaves remaining\n",
		e.FirstName,
		e.LastName,
		e.TotalLeaves-e.LeavesTaken,
	)
}

func Main() {
	e := Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

	var e2 Employee
	e2.LeavesRemaining()

	e3 := NewEmployee("Sam", "Adolf", 30, 20)
	e3.LeavesRemaining()
}
