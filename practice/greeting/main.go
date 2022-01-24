package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	greetingMessageEmpty := hello("")
	fmt.Println(aurora.Yellow(greetingMessageEmpty))

	greetingMessageJhon := hello("Jhon")
	fmt.Println(aurora.Yellow(greetingMessageJhon))
}

// simple syntax of testing
// import "testing"
// func TestAbc(t *testing.T) {
//     t.Error() // to indicate test failed
// }

//can create a unittesting cache with command
// go test ./unittesting(folder name) -c
//and run the test with
// ./unittesting.test -test.v
//it can save you time
//dont overwrite cache i there is no change in your file e.g. variable name change not gonna change the output or flow

//you can create a tessdata directory inside your package and put your file in there it wont be used during build
//and you can also use the test data from it
