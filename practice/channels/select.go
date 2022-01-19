// package main

// import (
// 	"fmt"
// 	"time"
// )

// var start time.Time

// func init() {
// 	start = time.Now()
// }

// func service1(c chan string) {
// 	time.Sleep(3 * time.Second)
// 	c <- "Hello from service 1"
// }

// func service2(c chan string) {
// 	time.Sleep(5 * time.Second)
// 	c <- "Hello from service 2"
// }

// func main() {
// 	fmt.Println("main() started", time.Since(start))

// 	chan1 := make(chan string)
// 	chan2 := make(chan string)

// 	go service1(chan1)
// 	go service2(chan2)

// 	select {
// 	case res := <-chan1:
// 		fmt.Println("Response from service 1", res, time.Since(start))
// 	case res := <-chan2:
// 		fmt.Println("Response from service 2", res, time.Since(start))
// 	}

// 	fmt.Println("main() stopped", time.Since(start))
// }

//without sleep

// package main

// import (
// 	"fmt"
// 	"time"
// )

// var start time.Time

// func init() {
// 	start = time.Now()
// }

// func service1(c chan string) {
// 	c <- "Hello from service 1"
// }

// func service2(c chan string) {
// 	c <- "Hello from service 2"
// }

// func main() {
// 	fmt.Println("main() started", time.Since(start))

// 	chan1 := make(chan string)
// 	chan2 := make(chan string)

// 	go service1(chan1)
// 	go service2(chan2)

// 	select {
// 	case res := <-chan1:
// 		fmt.Println("Response from service 1", res, time.Since(start))
// 	case res := <-chan2:
// 		fmt.Println("Response from service 2", res, time.Since(start))
// 	}

// 	fmt.Println("main() stopped", time.Since(start))
// }

//buffered select

package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
