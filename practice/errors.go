//example-1

// package main

// import "fmt"

// // create a struct
// type MyError struct{}

// // struct implements `Error` method
// func (myErr *MyError) Error() string {
// 	return "Something unexpected happend!"
// }

// func main() {

// 	// create error
// 	myErr := &MyError{}

// 	// print error message
// 	fmt.Println(myErr)
// }

//example-2

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	// create error
// 	myErr := errors.New("Something unexpected happend!")

// 	// print error type
// 	fmt.Printf("Type of myErr is %T \n", myErr)
// 	fmt.Printf("Value of myErr is %#v \n", myErr)
// }

//example-3

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// divide two number
// func Divide(a int, b int) (int, error) {

// 	// can not divide by `0`
// 	if b == 0 {
// 		return 0, errors.New("Can not devide by Zero!")
// 	} else {
// 		return (a / b), nil
// 	}
// }

// func main() {

// 	// divide 4 by 0
// 	if result, err := Divide(4, 0); err != nil {
// 		fmt.Println("Error occured: ", err)
// 	} else {
// 		fmt.Println("4/0 is", result)
// 	}
// }

//example-3

// package main

// import "fmt"

// // simple HTTP error
// type HttpError struct {
// 	status int
// 	method string
// }

// // HttpError implements Error method
// func (httpError *HttpError) Error() string {
// 	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
// 		httpError.method, httpError.status)
// }

// // mock function: sends HTTP request
// func GetUserEmail(userId int) (string, error) {

// 	// request failed, return an error
// 	return "", &HttpError{403, "GET"}
// }

// func main() {

// 	// get user email address
// 	if email, err := GetUserEmail(1); err != nil {
// 		fmt.Println(err) // print error

// 		// if user is unauthorized, perform session cleaning
// 		if errVal := err.(*HttpError); errVal.status == 403 {
// 			fmt.Println("Performing session cleaning...")
// 		}

// 	} else {
// 		// do something with the `email`
// 		fmt.Println("User email is", email)
// 	}
// }

//excample-4

// package main

// import "fmt"

// // network error
// type NetworkError struct{}

// func (e *NetworkError) Error() string {
// 	return "A network connection was aborted."
// }

// // file save failed error
// type FileSaveFailedError struct{}

// func (e *FileSaveFailedError) Error() string {
// 	return "The requested file could not be saved."
// }

// // a function that can return either of the above errors
// func saveFileToRemote() error {
// 	result := 2 // mock result of save operation

// 	if result == 1 {
// 		return &NetworkError{}
// 	} else if result == 2 {
// 		return &FileSaveFailedError{}
// 	} else {
// 		return nil
// 	}
// }

// func main() {
// 	// check type
// 	switch err := saveFileToRemote(); err.(type) {
// 	case nil:
// 		fmt.Println("File successfully saved.")
// 	case *NetworkError:
// 		fmt.Println("Network Error:", err)
// 	case *FileSaveFailedError:
// 		fmt.Println("File save Error:", err)
// 	}

// }

//example-5

// package main

// import "fmt"

// // simple user unauthorized error
// type UnauthorizedError struct {
// 	UserId        int
// 	OriginalError error
// }

// // add some context to the original error message
// func (httpErr *UnauthorizedError) Error() string {
// 	return fmt.Sprintf("User unauthorized Error: %v", httpErr.OriginalError)
// }

// // mock function call to validate user, returns error
// func validateUser(userId int) error {

// 	// mock general error from a function call: getSession(userId)
// 	err := fmt.Errorf("Session invalid for user id %d", userId)

// 	// return UnauthorizedError with original error
// 	return &UnauthorizedError{userId, err}
// }

// func main() {

// 	// validate user with id `1`
// 	err := validateUser(1)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("User is allowed to perform this action!")
// 	}

// }

//example-6

package main

import "fmt"

// simple user unauthorized error
type UnauthorizedError struct {
	UserId    int
	SessionId int
	error
}

// check if user is logged in
func (err *UnauthorizedError) isLoggedIn() bool {
	return err.SessionId != 0 // SessionId is 0 for non-logged in users
}

// mock function call to validate user, returns error
func validateUser(userId int) error {

	// mock general error from a function call: getSession(userId)
	err := fmt.Errorf("Session invalid for user id %d", userId)

	// return UnauthorizedError with original error
	return &UnauthorizedError{userId, 18783664, err}
}

func main() {

	// validate user with id `1`
	err := validateUser(1)

	if err != nil {

		// extract error object from `err` interface
		if errVal, ok := err.(*UnauthorizedError); ok {
			fmt.Println("Is user logged in:", errVal.isLoggedIn())
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println("User is allowed to perform this action!")
	}
}
