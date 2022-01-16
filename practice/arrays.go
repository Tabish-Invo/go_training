package main

import "fmt"

func main() {
	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println("array a => ", a)
	fmt.Println("elements => ", a[0], a[1], a[2])
}

// package main

// import "fmt"

// func main() {
// 	a := [3]int{1, 2}

// 	fmt.Println(a)
// }

//multiline intial value

// package main

// import "fmt"

// func main() {
// 	greetings := [4]string{
// 		"Good morning!",
// 		"Good afternoon!",
// 		"Good evening!",
// 		"Good night!", // must have comma
// 	}

// 	fmt.Println(greetings)
// }

//auto length

// package main

// import "fmt"

// func main() {
// 	greetings := [...]string{
// 		"Good morning!",
// 		"Good afternoon!",
// 		"Good evening!",
// 		"Good night!",
// 	}

// 	fmt.Println(greetings)
// }

//len of array

// package main

// import "fmt"

// func main() {
// 	greetings := [...]string{
// 		"Good morning!",
// 		"Good afternoon!",
// 		"Good evening!",
// 		"Good night!",
// 	}

// 	fmt.Println(len(greetings))
// }

//array comparision

// package main

// import "fmt"

// func main() {
// 	a := [3]int{1, 2, 3}
// 	b := [3]int{1, 3, 2}
// 	c := [3]int{1, 1, 1}
// 	d := [3]int{1, 2, 3}

// 	fmt.Println("a == b", a == b)
// 	fmt.Println("a == c", a == c)
// 	fmt.Println("a == d", a == d)
// }

//array iteratiom

// package main

// import "fmt"

// func main() {
// 	a := [...]int{1, 2, 3, 4, 5}

// 	for index := 0; index < len(a); index++ {
// 		fmt.Printf("a[%d] = %d\n", index, a[index])
// 	}
// }

//index,value with range

// package main

// import "fmt"

// func main() {
// 	a := [...]int{1, 2, 3, 4, 5}

// 	for index, value := range a {
// 		fmt.Printf("a[%d] = %d\n", index, value)
// 	}
// }

//multidimentional array

// package main

// import "fmt"

// func main() {
// 	a := [3][2]int{
// 		[2]int{1, 2},
// 		[2]int{3, 4},
// 	}

// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	a := [...][2]int{
// 		[...]int{1, 2},
// 		[...]int{3, 4},
// 		[...]int{5, 6},
// 	}

// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

// 	fmt.Printf("Array is %v and type of array element is %T", a, a[0])
// 	fmt.Println()
// }
