package main

import "fmt"

func main() {

	fmt.Println("experiment with struct in golang")

	type Employee struct {
		ID int
		Name string
		Address string
	//	DOB time.Time
		Position string
		Salary int
		ManagerID int
	}

	var dilbert Employee
	
	dilbert.Salary = -5000

	dilbert.ManagerID = 312

	fmt.Printf("%d", dilbert.ManagerID)
        Employeeid(312).Salary
}

func Employeeid(id int) *Employee {
}
