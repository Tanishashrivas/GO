package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string= "https://jsonplaceholder.typicode.com/posts/1";

func main() {
	// GET
	response, err := http.Get(url);

	checkNilErr(err)

	fmt.Printf("type of response is %T\n", response) // *http.Response: reference of original response
	// fmt.Println(response)

	defer response.Body.Close() // important to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)

	// fmt.Println("databytes:", databytes)
	// fmt.Printf("data: %s", databytes)
	fmt.Println("content:", string(databytes))

	// POST
}

func checkNilErr (err error) {
	if err != nil {
		panic(err)
	}
}