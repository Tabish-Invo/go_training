// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync"
// )

// func main() {

// 	wg := new(sync.WaitGroup)
// 	wg.Add(2)
// 	// create a default route handler
// 	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Fprint(res, "Hello: "+req.Host)
// 	})

// 	// create a goroutine
// 	go func() {
// 		// spawn an HTTP server in `other` goroutine
// 		log.Fatal(http.ListenAndServe(":9000", nil))
// 		wg.Done()
// 	}()

// 	// // spawn an HTTP server in `main` goroutine
// 	// log.Fatal(http.ListenAndServe(":9001", nil))

// 	go func() {
// 		// spawn an HTTP server in `other` goroutine
// 		log.Fatal(http.ListenAndServe(":9001", nil))
// 		wg.Done()
// 	}()

// 	wg.Wait()

// }

// //Example-2

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func createServer(add string, port int) *http.Server {

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
// 		fmt.Print(rw, "Hello "+add)
// 	})

// 	//create server
// 	server := http.Server{
// 		Addr:    fmt.Sprint(":%v", add),
// 		Handler: mux,
// 	}

// 	return &server
// }

// //close method

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {

// 	// create channels for safe exit
// 	exitSignal := make(chan interface{})

// 	// create server to run on port the 9000
// 	server := &http.Server{
// 		Addr:    ":9000",
// 		Handler: nil, // use `DefaultServeMux`
// 	}

// 	// close server after 3 seconds
// 	time.AfterFunc(10*time.Second, func() {
// 		fmt.Println("Close(): completed!", server.Close()) // close `server`
// 		close(exitSignal)                                  // close `exitSignal` channel
// 	})

// 	// launch server
// 	err := server.ListenAndServe()
// 	fmt.Println("ListenAndServe():", err)

// 	// listen to `exitSignal` channel
// 	<-exitSignal
// 	fmt.Println("Main(): exit complete!")
// }

// //Shtudown

package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	// create a `WaitGroup` for safe exit
	wg := new(sync.WaitGroup)
	wg.Add(2) // add `2` goroutines to finish

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `DefaultServeMux`
	}

	// register a cleanup function
	server.RegisterOnShutdown(func() {
		fmt.Println("RegisterOnShutdown(): completed!") // perform a cleanup
		wg.Done()                                       // WaitGroup done
	})

	// shutdown server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		err := server.Shutdown(context.Background()) // shutdown `server`
		fmt.Println("Shutdown(): completed!", err)
		wg.Done() // WaitGroup done
	})

	// launch server
	fmt.Println("ListenAndServe():", server.ListenAndServe())

	// listen for `wg` to complete
	wg.Wait()
	fmt.Println("Main(): exit complete!")
}
