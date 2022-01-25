// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	fs := http.FileServer(http.Dir("/Users/tabishamin/Desktop/Garbage"))

// 	fmt.Println("\nServer Started !\n\nNavigate To http://127.0.0.1:9000/ ")
// 	log.Fatal(http.ListenAndServe(":9000", fs))
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	// create file server handler
// 	fs := http.FileServer(http.Dir("/Users/tabishamin/Desktop/Garbage"))

// 	// handle `/` route
// 	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 		res.Header().Set("Content-Type", "text/html")
// 		fmt.Fprint(res, "<h1>Golang!</h1>")
// 	})

// 	// handle `/static` route
// 	// http.Handle("/static", fs)
// 	http.Handle("/static/", http.StripPrefix("/static", fs))

// 	// start HTTP server with `http.DefaultServeMux` handler
// 	log.Fatal(http.ListenAndServe(":9000", nil))
// }

// package main

// import (
// 	"log"
// 	"net/http"
// 	"path/filepath"
// )

// // temporary directory location
// var tmpDir = filepath.FromSlash("/Users/tabishamin/Desktop/Garbage/")

// func main() {

// 	// return a `.pdf` file for `/pdf` route
// 	http.HandleFunc("/pdf", func(res http.ResponseWriter, req *http.Request) {
// 		http.ServeFile(res, req, filepath.Join(tmpDir, "/files/Tabish-Referral_Contacts.pdf"))
// 	})

// 	// start HTTP server with `http.DefaultServeMux` handler
// 	log.Fatal(http.ListenAndServe(":9000", nil))

// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

// temporary directory location
var tmpDir = filepath.FromSlash("/Users/tabishamin/Desktop/Garbage")

func main() {

	// default route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// return a `.html` file for `/index.html` route
	http.HandleFunc("/index.html", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, filepath.Join(tmpDir, "/index.html"))
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))

}

// exePath, err := os.Executable()   => /usr/share/program.exe
// exeDir := filepath.Dir(exePath) => /usr/share/
// os.Getwd()
