// Go program to illustrate the
// use of simple for loop
package main

import "fmt"

// Main function
func main() {

	// for loop
	// This loop starts when i = 0
	// executes till i<4 condition is true
	// post statement is i++
	for i := 0; i < 4; i++ {
		fmt.Printf("Hello World!\n")
	}

}

// Go program to illustrate the
// use of an infinite loop
// package main

// import "fmt"

// // Main function
// func main() {

// 	// Infinite loop
// 	for {
// 	fmt.Printf("Hello World!\n")
// 	}

// }

// Go program to illustrate the
// the for loop as while Loop
// package main

// import "fmt"

// // Main function
// func main() {

// 	// while loop
// 	// for loop executes till
// 	// i < 3 condition is true
// 	i:= 0
// 	for i < 3 {
// 	i += 2
// 	}
// fmt.Println(i)
// }

// Go program to illustrate the
// use of simple range loop
// package main

// import "fmt"

// // Main function
// func main() {

// 	// Here rvariable is a array
// 	rvariable:= []string{"GFG", "Geeks", "GeeksforGeeks"}

// 	// i and j stores the value of rvariable
// 	// i store index number of individual string and
// 	// j store individual string of the given array
// 	for i, j:= range rvariable {
// 	fmt.Println(i, j)
// 	}

// }

// Go program to illustrate the
// use of simple range loop
// package main

// import "fmt"

// // Main function
// func main() {

// 	// Here rvariable is a array
// 	rvariable:= []string{"GFG", "Geeks", "GeeksforGeeks"}

// 	// i and j stores the value of rvariable
// 	// i store index number of individual string and
// 	// j store individual string of the given array
// 	for i, j:= range rvariable {
// 	fmt.Println(i, j)
// 	}

// }

//Without all statments

// package main

// import "fmt"

// func main() {
// 	i := 1
// 	for {
// 		fmt.Printf("Current number is %d \n", i)

// 		if i == 6 {
// 			break
// 		}
// 		i++
// 	}
// }

//break,continue
