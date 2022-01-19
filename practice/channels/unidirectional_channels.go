package main

import "fmt"

// func main() {
// 	roc := make(<-chan int)
// 	soc := make(chan<- int)

// 	fmt.Printf("Data type of roc is `%T`\n", roc)
// 	fmt.Printf("Data type of soc is `%T\n", soc)
// }

func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
