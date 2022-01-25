package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello")
	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " World! ")
	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
}

func main() {
	// // create a new handler
	// handler := HttpHandler{}

	fmt.Println("\nServer Started !\n\nNavigate To http://127.0.0.1:9000/ ")

	// // listen and serve
	// http.ListenAndServe(":9000", handler)

	//with mux

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(res, "Home Page")
	// })
	// mux.HandleFunc("/golang", func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(res, "Golang Page")
	// })
	// http.ListenAndServe(":9000", mux)

	//without mux

	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(res, "Home Page")
	// })

	// // handle `/hello/golang` route to `http.DefaultServeMux`
	// http.HandleFunc("/golang", func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(res, "Golang Page")
	// })

	// http.ListenAndServe(":9000", nil)

	//set header

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2020")

		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400

		// respond with a JSON string
		fmt.Fprint(res, `{"status":"FAILURE"}`)
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))
}
