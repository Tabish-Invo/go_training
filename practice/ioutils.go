// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// const tempDir = "/Users/tabishamin/Desktop/Garbage"

// func main() {

// 	files, _ := ioutil.ReadDir(tempDir)

// 	for _, file := range files {

// 		fmt.Printf("Name: %v, Size: %vkb, Mode: %v, IsDir: %v\n", file.Name(), file.Size(), file.Mode(), file.IsDir())
// 	}

// }

//GLOB

// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {

// 	// search for `.pdf` files
// 	pdfFiles, _ := filepath.Glob("/Users/tabishamin/Desktop/Garbage/**/*.pdf")

// 	// list information of each file
// 	for _, pdfFile := range pdfFiles {

// 		file, _ := os.Stat(pdfFile)

// 		fmt.Printf("Name: %v, Size: %v kb, Mode: %v, IsDir: %v\n",
// 			file.Name(),
// 			file.Size()/1000,
// 			file.Mode(),
// 			file.IsDir(),
// 		)
// 	}

// }

//read file content

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {

// 	html, _ := ioutil.ReadFile("/Users/tabishamin/Desktop/Garbage/index.html")

// 	// print raw bytes
// 	fmt.Println("Raw: \n", html)

// 	// print string representation
// 	fmt.Println("String: \n", string(html))

// }

// write file content

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// temporary directory
const tmpDir = "/Users/tabishamin/Desktop/Garbage"

func main() {

	// welcome message content
	welcomeData := []byte("File Created By Code.")

	// get `welcome.txt` file path
	path := filepath.Join(tmpDir, "/welcome.txt")

	// write content to `welcome.txt` file
	err := ioutil.WriteFile(path, welcomeData, 0777)

	// log error
	if err != nil {
		fmt.Println(err)
	}

}
