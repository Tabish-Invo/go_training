// package main

// import "fmt"

// func defFunc() {

// 	fmt.Println("defFucn() started..!")

// 	if r := recover(); r != nil {
// 		fmt.Println("Program is unpacking with value ", r)
// 	}

// 	fmt.Println("defFucn() ended..!")
// }

// func norm() {

// 	fmt.Println("norm() started..!")

// 	defer defFunc()
// 	panic("Help")

// 	fmt.Println("norm() ended..!")
// }

// func main() {

// 	fmt.Println("main() started..!")

// 	norm()

// 	fmt.Println("main() ended..!")
// }

// // example-2

// package main

// import "fmt"

// func defBar() {
// 	fmt.Println("defBar() started")

// 	if r := recover(); r != nil {
// 		fmt.Println("WOHA! Program is panicking with value", r)
// 	}

// 	fmt.Println("defBar() done")
// }

// func defFoo() {
// 	fmt.Println("defFoo() started")

// 	defer defBar() // defer call

// 	fmt.Println("defFoo() done")
// }

// func normMain() {

// 	fmt.Println("normMain() started")

// 	defer defFoo() // defer defFoo call

// 	panic("HELP") // panic here

// 	fmt.Println("normMain() done")
// }

// func defMain() {

// 	fmt.Println("defMain() started")

// 	fmt.Println("defMain() done")
// }

// func main() {
// 	fmt.Println("main() started")

// 	defer defMain() // defer call

// 	normMain() // normal call

// 	fmt.Println("main() done")
// }

// //example-3

package main

import "fmt"

// access an element of a slice by index
func accessElement(a []int, index int) (v int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered!")
			v = a[len(a)-1] // set v to the last value
		}
	}()

	v = a[index]

	return
}

func main() {
	a := []int{1, 2, 3}

	// access 3rd element
	fmt.Println(accessElement(a, 2))

	// access 4th element (out of bound)
	fmt.Println(accessElement(a, 3))
}
