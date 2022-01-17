//save hex values
// package main

// import "fmt"

// func main() {
// 	a := 0x00
// 	b := 0x0A
// 	c := 0xFF

// 	fmt.Printf("variable a of type %T with value %v in hex is %X\n", a, a, a)
// 	fmt.Printf("variable b of type %T with value %v in hex is %X\n", b, b, b)
// 	fmt.Printf("variable c of type %T with value %v in hex is %X\n", c, c, c)

// 	fmt.Println("&a = ", &a)
// 	fmt.Println("&b = ", &b)
// 	fmt.Println("&c = ", &c)
// }

//sipmle pointer
package main

import "fmt"

func changevalue(p *int) {
	*p = 123
}

func changevalue_arr(p *[3]int) {
	(*p)[0] += 1
	p[1] *= 2
	p[2] -= 3
}

func main() {
	a := 1
	var pa *int //not compulsory
	pa = &a

	fmt.Printf("Pointer pa of type %T with value %v\n", pa, pa)

	//derefrence a pointer

	fmt.Printf("Value of %v is %v\n", pa, *pa)

	//chnage value using pointer

	*pa = 2

	fmt.Printf("Value of a is %v\n", a)
	fmt.Printf("Value of pa is %v\n", *pa)

	//new function

	new_p := new(int)

	fmt.Printf("Data at %v is %v\n", new_p, *new_p)

	//pass pointer to function

	ab := 111
	p_ab := &ab

	changevalue(p_ab)
	//changevalue(&ab)

	fmt.Printf("ab = %v\n", ab)

	//pass array as pointer to function

	arr := [3]int{3, 2, 7}
	changevalue_arr(&arr)

	fmt.Printf("arr = %v\n", arr)

}
