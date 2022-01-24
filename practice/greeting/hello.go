package main

import "fmt"

//return hello greetings

func hello(str string) string {

	if len(str) == 0 {
		return "Hello Dude!"
	} else {
		return fmt.Sprintf("Hello %v!", str)
	}
}
