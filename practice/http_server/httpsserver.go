package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//basic go server (Go’s standard server implementation)

	// http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ) {
	// 	fmt.Fprint( res, "Hello World!" )
	// } )

	// log.Fatal( http.ListenAndServeTLS( ":9000", "localhost.crt", "localhost.key", nil ) )

	//custom go server

	// s := &http.Server{
	// 	Addr:    ":9000",
	// 	Handler: nil,
	// }

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello World!")
	// })

	// log.Fatal(s.ListenAndServeTLS("localhost.crt", "localhost.key"))

	//TLS Certificates

	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")

	s := &http.Server{
		Addr:    ":9000",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World TSLConfig!")
	})

	log.Fatal(s.ListenAndServeTLS("", ""))

}
