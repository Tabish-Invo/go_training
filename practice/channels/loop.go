package main

import (
	"fmt"
	"time"
)

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		fmt.Println("In Square. Value of i == ", i)
		time.Sleep(time.Millisecond * 10)
		c <- i * i
	}

	close(c)
}

func main() {

	fmt.Println("main() started")

	c := make(chan int)

	go squares(c)

	// for {

	// 	value, ok := <-c
	// 	if ok == false {
	// 		fmt.Println(value, ok, "<----Loop Broken")
	// 		break
	// 	} else {
	// 		fmt.Println(value, ok)
	// 	}
	// }

	for value := range c {

		fmt.Println(value)
	}

	fmt.Println("main() stopped")
}
