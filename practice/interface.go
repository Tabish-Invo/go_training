package main

import (
	"fmt"
	"math"
	"strings"
)

type MyString string

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Object1 interface {
	Area() float64
}

type Object2 interface {
	Volume() float64
}

type Cube struct {
	side float64
}

type Skin interface {
	Color() float64
}

type Material interface {
	Object1
	Object2
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width * r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius * 2
}

func (cb Cube) Area() float64 {
	return 6 * (cb.side * cb.side)
}

func (cb Cube) Volume() float64 {
	return cb.side * cb.side * cb.side
}

func explain(i interface{}) {
	fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func explain_switch(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I Stored String which is : ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("I Stored Integer which is : ", i)
	default:
		fmt.Println("I Stored Something else which is :", i)
	}
}

func main() {

	var s Shape
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)

	var s1 Shape = Rect{10, 3}

	fmt.Printf("Type of s1 is %T\n", s1)
	fmt.Printf("Value of s1 is %v\n", s1)
	fmt.Printf("Area of s1 is %0.2f\n", s1.Area())

	var c1 Shape = Circle{10}

	fmt.Printf("Type of c1 is %T\n", c1)
	fmt.Printf("Value of c1 is %v\n", c1)
	fmt.Printf("Area of c1 is %0.2f\n", c1.Area())

	//In order to successfully implement an interface, you need to implement all the methods declared by the interface with exact signatures.

	//Empty Interface
	ms := MyString("Hello World!")
	r := Rect{5.5, 4.5}
	explain(ms)
	explain(r)

	c := Cube{3}
	var obj1 Object1 = c
	var obj2 Object2 = c

	fmt.Println("Area of obj1 of interface type Object1 is ", obj1.Area())
	fmt.Println("Volume of obj2 of interface type Object2 is ", obj2.Volume())

	//Type assertion

	var obj3 Object1 = c
	obj4 := obj3.(Cube)

	fmt.Println("Area of Obj4 of type Cube is ", obj4.Area())
	fmt.Println("Volume of Obj4 of type Cube is ", obj4.Volume())

	// check valid type assertion
	//value, ok := i.(Type)

	var s11 Object1 = Cube{3}
	value1, ok1 := s11.(Object2)
	fmt.Printf("dynamic value of Object1 's' with value %v implements interface Object2? %v\n", value1, ok1)
	value2, ok2 := s.(Skin)
	fmt.Printf("dynamic value of Object1 's' with value %v implements interface Skin? %v\n", value2, ok2)

	//Type switch

	explain_switch("Hello")
	explain_switch(123)
	explain_switch(true)

	//Embedding interfaces

	var mat Material = c
	var o1 Object1 = c
	var o2 Object2 = c

	fmt.Printf("dynamic type and value of interface mat of static type Material is'%T' and '%v'\n", mat, mat)
	fmt.Printf("dynamic type and value of interface o1 of static type Object1 is'%T' and '%v'\n", o1, o1)
	fmt.Printf("dynamic type and value of interface o2 of static type Object2 is'%T' and '%v'\n", o2, o2)

}
