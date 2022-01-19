package main

import "fmt"

func greet(c chan string) {
	<-c //for Jhon
	<-c //for Mike
}

func main() {
	fmt.Println("main started()")
	c := make(chan string)

	go greet(c)

	c <- "Jhon"

	close(c) //clossing channel

	c <- "Mike"

	fmt.Println("main stopped()")

}
