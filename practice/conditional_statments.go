// Go program to illustrate the
// use of if..else..if ladder
package main

import "fmt"

func main() {

	// taking a local variable
	var v1 int = 700

	// checking the condition
	if v1 == 100 {

		// if condition is true then
		// display the following */
		fmt.Printf("Value of v1 is 100\n")

	} else if v1 == 200 {

		fmt.Printf("Value of a is 20\n")

	} else if v1 == 300 {

		fmt.Printf("Value of a is 300\n")

	} else {

		// if none of the conditions is true
		fmt.Printf("None of the values is matching\n")
	}
}

// Go program to illustrate the
// concept of Type switch
// statement
// package main

// import "fmt"

// func main() {
// 	var value interface{}
// 	switch q:= value.(type) {
// 	case bool:
// 	fmt.Println("value is of boolean type")
// 	case float64:
// 	fmt.Println("value is of float64 type")
// 	case int:
// 	fmt.Println("value is of int type")
// 	default:
// 	fmt.Printf("value is of type: %T", q)

// }
// }

//Go program to illustrate the
//concept of if-else with intial statment

// package main

// import "fmt"

// func main() {
// 	if fruit := "banana"; fruit == "mango" {
// 		fmt.Println("fruit is mango")
// 	} else if fruit == "orange" {
// 		fmt.Println("fruit is orange")
// 	} else if fruit == "banana" {
// 		fmt.Println("fruit is banana")
// 	} else {
// 		fmt.Println("I don't know which fruit this is")
// 	}

// 	// fruit variable is unavailable here
// 	// fmt.Println(fruit)
// }

//Multiple match in switch
// import "fmt"

// func main() {
// 	letter := "i"

// 	switch letter {
// 	case "a", "e", "i", "o", "u":
// 		fmt.Println("Letter is a vovel.")
// 	default:
// 		fmt.Println("Letter is not a vovel.")
// 	}
// }

//Expressionless Switch Statment

// package main

// import "fmt"

// func main() {
// 	number := 20

// 	switch {
// 	case number <= 5:
// 		fmt.Println("number is less than or equal to 5")
// 	case number > 5:
// 		fmt.Println("number is greater than 5")
// 	case number > 10:
// 		fmt.Println("number is greater than 10")
// 	case number > 15:
// 		fmt.Println("number is greater than 15")
// 	}
// }

//Fallthrough Statment

// package main

// import "fmt"

// func main() {
// 	switch number := 20; {
// 	case number <= 5:
// 		fmt.Println("number is less than or equal to 5")
// 		fallthrough
// 	case number > 5:
// 		fmt.Println("number is greater than 5")
// 		fallthrough
// 	case number > 10:
// 		fmt.Println("number is greater than 10")
// 		fallthrough
// 	case number > 15:
// 		fmt.Println("number is greater than 15")
// 	}
// }
