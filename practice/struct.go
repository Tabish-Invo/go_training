//simple struct

package main

import (
	"fmt"
	"golang/org"
)

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

//anonymous fields struct

type Data struct {
	string
	int
	bool
}

//nested struct

type Salary struct {
	basic     int
	insurance int
	allowance int
}

type Emp struct {
	firstName, lastName string
	Salary
	bool
}

//nested interface

type Salaried interface {
	getsalary() int
}

func (s Salary) getsalary() int {
	return s.basic + s.allowance + s.insurance
}

type Employee_ struct {
	firstName, lastName string
	salary              Salaried
	fullTime            bool
}

//function fields

type FullNameType func(string, string) string

type Emp_ struct {
	FirstName, LastName string
	FullName            FullNameType
}

func main() {
	var ross Employee
	fmt.Println(ross)

	ross.firstName = "Mike"
	ross.lastName = "Ross"
	ross.salary = 100000
	ross.fullTime = true

	fmt.Println(ross)

	tabish := Employee{
		fullTime:  true,
		firstName: "Tabish",
		lastName:  "Amin",
		salary:    200000,
	}

	fmt.Println("Tabish's Data :")
	fmt.Println("First Name = ", tabish.firstName)
	fmt.Println("Last Name = ", tabish.lastName)
	fmt.Println("Full Time = ", tabish.fullTime)
	fmt.Println("Salary = ", tabish.salary)

	//anonymous struct

	monica := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:    1200,
	}

	fmt.Println(monica)

	//pointer to a struct

	ross_ := &Employee{
		firstName: "ross",
		lastName:  "Bing",
		salary:    1200,
		fullTime:  false,
	}

	fmt.Println("firstName", (*ross_).firstName)

	// Anonymous Fields

	anon := Data{"Monday", 1, true}
	anon.bool = true

	fmt.Println("Anonymous Struct = ", anon)

	//Nested struct

	geller := Emp{
		firstName: "Ross",
		lastName:  "Geller",
		bool:      true,
		Salary:    Salary{1100, 50, 50},
	}
	fmt.Println(geller)
	fmt.Println("Geller's Last Name = ", geller.lastName)
	fmt.Println("Geller's basic salary", geller.basic)

	// Nested interface

	new_ross := Employee_{
		firstName: "New",
		lastName:  "Ross",
		salary:    Salary{1100, 500, 200},
	}

	fmt.Println("New Ross Salary is = ", new_ross.salary.getsalary())

	// Exported fields

	emp_obj := org.Employee_pkg{
		FirstName: "Tabish",
		LastName:  "Amin",
	}

	fmt.Println(emp_obj)

	//Function fields

	e := Emp_{
		FirstName: "Ali",
		LastName:  "Raza",
		FullName: func(firstName, lastName string) string {
			return firstName + " " + lastName
		},
	}

	fmt.Printf(e.FullName(e.FirstName, e.LastName))

}
