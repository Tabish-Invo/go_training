package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	// // make a sample HTTP GET request
	// res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, _ := ioutil.ReadAll(res.Body)

	// res.Body.Close()

	// fmt.Printf("%s\n", data)

	// fmt.Println(res.Header)

	// contentType := res.Header.Get("Content-Type")
	// fmt.Println("Content-Type", contentType)

	//download image to understnad GET request

	// res, err := http.Get("https://bit.ly/2IRnmVm")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if res.Header.Get("Content-Type") == "image/jpeg" {
	// 	ioutil.WriteFile("download_img.jpeg", data, 0777)
	// 	fmt.Printf("Image Saved\n")
	// } else {
	// 	log.Fatal("Error : Response in not an image")
	// }

	//Post Request

	// requestBody := strings.NewReader(`
	// 	{
	// 		"name":"test",
	// 		"salary":"123",
	// 		"age":"23"
	// 	}
	// `)

	// res, err := http.Post(
	// 	"http://dummy.restapiexample.com/api/v1/create",
	// 	"application/json; charset=UTF-8",
	// 	requestBody,
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, _ := ioutil.ReadAll(res.Body)

	// res.Body.Close()

	// requestContentType := res.Request.Header.Get("Content-Type")
	// fmt.Println("Request content-type:", requestContentType)

	// fmt.Printf("%s\n", data)

	//example-2 post

	// url, _ := url.Parse("http://dummy.restapiexample.com/api/v1/create")
	// reqBody := ioutil.NopCloser(strings.NewReader(`{
	// 		"name":"test",
	// 		"salary":"123",
	// 		"age":"23"
	// }`))

	// req := &http.Request{
	// 	Method: "POST",
	// 	URL:    url,
	// 	Header: map[string][]string{
	// 		"Content-Type": {"application/json;charset=UFT-8"},
	// 	},
	// 	Body: reqBody,
	// }

	// res, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, _ := ioutil.ReadAll(res.Body)

	// res.Body.Close()

	// fmt.Printf("Status Code : %d\n", res.StatusCode)
	// fmt.Printf("Body : %s\n", data)

	// error
	client := &http.Client{
		Timeout: 100 * time.Microsecond,
	}

	res, err := client.Get("https://jsonplaceholder.typicode.com/users/1")

	// check for response error
	if err != nil {

		urlErr := err.(*url.Error)

		if urlErr.Timeout() {
			fmt.Println("Error occurred due to a timeout.")
		}

		log.Fatal("Error:", err)
	} else {
		fmt.Println("Success: status-code", res.StatusCode)
	}

}
