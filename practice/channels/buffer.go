package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	fmt.Println("calling 1")
	c <- 1
	fmt.Println("calling 2")
	c <- 2
	fmt.Println("calling 3")
	c <- 3

	//blocking statment
	fmt.Println("calling 4")
	c <- 4

	fmt.Println("main() stopped")
}
