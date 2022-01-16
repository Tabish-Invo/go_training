package main

import "fmt"

func add(a, b int) int64 {
	return int64(a + b)
}

func main() {
	result := add(1, 5)
	fmt.Println(result)
}

//multiple return values

// package main

// import "fmt"

// func addMult(a, b int) (int, int) {
// 	return a + b, a * b
// }

// func main() {
// 	addRes, multRes := addMult(2, 5)
// 	fmt.Println(addRes, multRes)
// }

//named return values

// package main

// import "fmt"

// func addMult(a, b int) (add int, mul int) {
// 	add = a + b
// 	mul = a * b

// 	return // necessary
// }

// func main() {
// 	addRes, multRes := addMult(2, 5)
// 	fmt.Println(addRes, multRes)
// }

//recursive function

// package main

// import "fmt"

// // n! = n×(n-1)! where n >0
// func getFactorial(num int) int {
// 	if num > 1 {
// 		return num * getFactorial(num-1)
// 	}

// 	return 1 // 1! == 1
// }

// func main() {
// 	f := getFactorial(4)
// 	fmt.Println(f)
// }

//defer

// package main

// import "fmt"

// func endTime(timestamp string) {
// 	fmt.Println("Program ended at", timestamp)
// }

// func main() {
// 	time := "1 PM"

// 	defer endTime(time)

// 	time = "2 PM"

// 	fmt.Println("doing something")
// 	fmt.Println("main finished")
// 	fmt.Println("time is", time)
// }

// package main

// import "fmt"

// func greet(message string) {
// 	fmt.Println("greeting: ", message)
// }

// func main() {
// 	fmt.Println("Call one")

// 	defer greet("Greet one")

// 	fmt.Println("Call two")

// 	defer greet("Greet two")

// 	fmt.Println("Call three")

// 	defer greet("Greet three")
// }

//function as a type

// package main

// import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func subtract(a int, b int) int {
// 	return a - b
// }

// func main() {
// 	fmt.Printf("Type of function add is			%T\n", add)
// 	fmt.Printf("Type of function subtract is		%T\n", subtract)
// }

// package main

// import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func subtract(a int, b int) int {
// 	return a - b
// }

// func calc(a int, b int, f func(int, int) int) int {
// 	r := f(a, b)
// 	return r
// }

// func main() {
// 	addResult := calc(5, 3, add)
// 	subResult := calc(5, 3, subtract)
// 	fmt.Println("5+3 =", addResult)
// 	fmt.Println("5-3 =", subResult)
// }

//function as value

// package main

// import "fmt"

// var add = func(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	fmt.Println("5+3 =", add(5, 3))
// }

//Immediately-invoked function expression (IIFE)

// package main

// import "fmt"

// func main() {
// 	sum := func(a int, b int) int {
// 		return a + b
// 	}(3, 5)

// 	fmt.Println("5+3 =", sum)
// }
