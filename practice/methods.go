package main

import (
	"fmt"
	"math"
	"strings"
)

type MyString string

//basic method

type Employee struct {
	FirstName, LastName string
}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

// same name methods

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//Pointer receivers

type Employee_ struct {
	name   string
	salary int
}

func (e *Employee_) changeName(newName string) {
	e.name = newName
}

//Methods on nested struct

type Contact struct {
	phone, address string
}
type Emp struct {
	name    string
	salary  int
	contact Contact
}

func (c *Contact) changePhone(newPhone string) {
	c.phone = newPhone
}

//Methods on non-struct type

func (s MyString) toUpperCase() string {
	normalString := string(s)
	return strings.ToUpper(normalString)
}

func main() {

	//basic methods
	e := Employee{
		FirstName: "Mike",
		LastName:  "Ross",
	}

	fmt.Println(e.fullName())

	//same name methods
	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Area())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.Area())

	//Pointer receivers

	e_ := Employee_{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e_)
	// change name
	e_.changeName("Monica Geller")
	// e after name change
	fmt.Println("e after name change =", e_)

	//Methods on nested struct

	e1 := Emp{
		name:   "Ross Geller",
		salary: 1200,
		contact: Contact{
			phone:   "011 8080 8080",
			address: "New Delhi, India",
		},
	}
	// e before phone change
	fmt.Println("e1 before phone change =", e1)
	// change phone
	e1.contact.changePhone("011 1010 1222")
	// e after phone change
	fmt.Println("e1 after phone change =", e1)

	//Methods can accept both pointer and value

	//Methods on non-struct type
	str := MyString("Hello World")
	fmt.Println(str.toUpperCase())

}
