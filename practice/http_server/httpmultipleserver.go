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

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// create channels for safe exit
	exitSignal := make(chan interface{})

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `DefaultServeMux`
	}

	// close server after 3 seconds
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("Close(): completed!", server.Close()) // close `server`
		close(exitSignal)                                  // close `exitSignal` channel
	})

	// launch server
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)

	// listen to `exitSignal` channel
	<-exitSignal
	fmt.Println("Main(): exit complete!")
}
