package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println(m)
	fmt.Println("m == nil", m == nil)
}

//empty map

// package main

// import "fmt"

// func main() {
// 	age := make(map[string]int)
// 	age["mina"] = 28
// 	age["john"] = 32
// 	age["mike"] = 55
// 	fmt.Println("age of john", age["john"])
// }

//initializing a map

// package main

// import "fmt"

// func main() {
// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}

// 	fmt.Println(age)
// }

//accessing map data

// package main

// import "fmt"

// func main() {
// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}
// 	fmt.Println(age["mina"])
// 	fmt.Println(age["jessy"])
// }

//len of map

// package main

// import "fmt"

// func main() {
// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}
// 	fmt.Println("len(age) =", len(age))
// }

//delete map

// package main

// import "fmt"

// func main() {
// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}

// 	delete(age, "john")
// 	delete(age, "jessy")

// 	fmt.Println(age)
// }

//map iteration

// package main

// import "fmt"

// func main() {
// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}

// 	for key, value := range age {
// 		fmt.Println(key, "=>", value)
// 	}
// }

//maps with other data types

// package main

// import "fmt"

// func main() {
// 	age := map[bool]string{
// 		true:  "YES",
// 		false: "NO",
// 	}

// 	for key, value := range age {
// 		fmt.Println(key, "=>", value)
// 	}
// }

//maps are refrenced types

// package main

// import "fmt"

// func main() {
// 	var ages map[string]int

// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}

// 	ages = age

// 	delete(ages, "john")

// 	fmt.Println("age", age)
// 	fmt.Println("ages", ages)
// }
