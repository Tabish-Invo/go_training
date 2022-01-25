// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func printNumbers() {
// 	counter := 1

// 	for {
// 		time.Sleep(100 * time.Millisecond)
// 		counter++
// 	}
// }

// func main() {
// 	go printNumbers()

// 	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
// 	time.Sleep(time.Second)

// 	fmt.Println("After: active goroutines", runtime.NumGoroutine())
// 	fmt.Println("Program exited")
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// var signal = make(chan bool)

// func printNumbers() {
// 	counter := 1

// 	for {
// 		select {
// 		case <-signal:
// 			return
// 		default:
// 			time.Sleep(100 * time.Millisecond)
// 			counter++
// 		}
// 	}
// }

// func main() {
// 	go printNumbers()

// 	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
// 	time.Sleep(time.Second)

// 	signal <- true

// 	fmt.Println("After: active goroutines", runtime.NumGoroutine())
// 	fmt.Println("Program exited")
// }

// //main_cancel

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan struct{})

// 	run := func(ctx context.Context) {
// 		n := 1
// 		for {
// 			select {
// 			case <-ctx.Done(): // 2. "ctx" is cancelled, we close "ch"
// 				fmt.Println("exiting")
// 				close(ch)
// 				return // returning not to leak the goroutine
// 			default:
// 				time.Sleep(time.Millisecond * 300)
// 				fmt.Println(n)
// 				n++
// 			}
// 		}
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		fmt.Println("goodbye")
// 		cancel() // 1. cancels "ctx"
// 	}()

// 	go run(ctx)

// 	fmt.Println("waiting to cancel...")

// 	<-ch // 3. "ch" is closed, we exit

// 	fmt.Println("bye")
// }

// //main_deadline

package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
