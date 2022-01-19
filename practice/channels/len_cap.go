package main

import "fmt"

// func main() {

// 	c := make(chan int, 3)

// 	c <- 1
// 	c <- 2

// 	fmt.Printf("The Length of Channel c is %v and the Capacity is %v", len(c), cap(c))
// }

// If you are wondering, why the above program runs well and deadlock error was not thrown.
// This is because, as channel capacity is 3 and only 2 values are available in the buffer,
// Go did not try to schedule another goroutine by blocking main goroutine execution.
// You can simply read these value in the main goroutine if you want because even if the buffer is not full,
// that doesn’t prevent you to read values from the channel.

func sender(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocking statmnet
	fmt.Println("End Sender")
	close(c)
	fmt.Println("End Sender 2")
}

func main() {

	c := make(chan int, 3)
	go sender(c)
	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))
	// read values from c (blocked here)
	for val := range c {
		fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
	}
	fmt.Println("End Main")

}
