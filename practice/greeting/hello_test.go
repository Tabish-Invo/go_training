package main

import (
	"testing"
)

// func TestHello(t *testing.T) {

// 	emptyResult := hello("")
// 	if emptyResult != "Hello Dude!" {
// 		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
// 	} else {
// 		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Dude!", emptyResult)
// 	}

// 	result := hello("Mike")
// 	if result != "Hello Mike!" {
// 		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Mike!", result)
// 	} else {
// 		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Mike!", result)
// 	}
// }

func TestHelloEmptyArgs(t *testing.T) {

	emptyResult := hello("")
	if emptyResult != "Hello Dude!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
	} else {
		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Dude!", emptyResult)
	}
}

func TestHelloVAlidArgs(t *testing.T) {

	result := hello("Mike")
	if result != "Hello Mike!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Mike!", result)
	} else {
		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Mike!", result)
	}
}

//we can install go get -u github.com/rakyll/gotest
//gotest -v -run TestU //this creates a like function search

// go test -cover shows us coverage of our test run
//go test -coverprofile=cover.txt writes cover profile in cover.txt
//we can use go tool cover -html=cover.txt -o cover.html
