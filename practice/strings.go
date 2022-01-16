//simple string

package main

import "fmt"

func main() {
	var s string

	s = "Hello World"

	fmt.Println(s)
}

//length of string

// package main

// import "fmt"

// func main() {
// 	s := "Hello World"

// 	fmt.Println(len(s))
// }

//loop ove string

// package main

// import "fmt"

// func main() {
// 	s := "Hello World"

// 	for i := 0; i < len(s); i++ {
// 		fmt.Print(s[i], " ")
// 	}

// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	s := "Hello World"
// 	fmt.Println("len(s)", len(s))

// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c ", s[i])
// 	}

// 	fmt.Println("")

// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%v ", s[i])
// 	}

// 	fmt.Println("")

// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}

// 	fmt.Println("")
//         for i := 0; i < len(s); i++ {
// 		fmt.Printf("%T ", s[i])
// 	}

// 	fmt.Println("")
// }

//Latin alphabets
// package main

// import "fmt"

// func main() {
//     s := "Hellõ World"
// 	fmt.Println("len(s)", len(s))

// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c ", s[i])
// 	}

// 	fmt.Println("")

// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%v ", s[i])
// 	}

// 	fmt.Println("")

// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}

// 	fmt.Println("")
// }

//rune

// package main

// import "fmt"

// func main() {
// 	s := "Hellõ World"
// 	r := []rune(s)

// 	fmt.Println("len(r)", len(r))

// 	for i := 0; i < len(r); i++ {
// 		fmt.Printf("%c ", r[i])
// 	}
// 	fmt.Println("")

// 	for i := 0; i < len(r); i++ {
// 		fmt.Printf("%v ", r[i])
// 	}
// 	fmt.Println("")

// 	for i := 0; i < len(r); i++ {
// 		fmt.Printf("%x ", r[i])
// 	}
// 	fmt.Println("")

// 	for i := 0; i < len(r); i++ {
// 		fmt.Printf("%T ", r[i])
// 	}
// 	fmt.Println("")
// }

//loop

// package main

// import "fmt"

// func main() {
// 	s := "Hellõ World"

// 	for index, char := range s {
// 		fmt.Printf("character at index %d is %c\n", index, char)
// 	}
// }

//Strings are miutable

//String Litrals using backticks

// package main

// import "fmt"

// func main() {
// 	s := `Hello,\n
// 		My Big Blue
// 	"World"!`

// 	fmt.Println(s)
// }

//Character Comparision

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
// 	fmt.Printf("value of character b is %v of type %T\n", 'b', 'b')
// 	fmt.Println("hence 'b' > 'a' is", 'b' > 'a')
// }
