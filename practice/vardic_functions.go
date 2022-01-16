package main

import "fmt"

func getMultiples(factor int, args ...int) []int {
	multiples := make([]int, len(args))

	for index, val := range args {
		multiples[index] = val * factor
	}

	return multiples
}

func main() {
	s := []int{10, 20, 30}

	// get multiples of 2, pass parameters from slice using `unpack` operator
	mult1 := getMultiples(2, s...)

	// get multiples of 3
	mult2 := getMultiples(3, 1, 2, 3, 4)

	fmt.Println(mult1)
	fmt.Println(mult2)
}

// package main

// import "fmt"

// func getMultiples(factor int, args ...int) []int {
// 	for index, val := range args {
// 		args[index] = val * factor
// 	}

// 	return args
// }

// func main() {
// 	s := []int{10, 20, 30}

// 	mult1 := getMultiples(2, s...)

// 	fmt.Println(s, mult1)
// }
