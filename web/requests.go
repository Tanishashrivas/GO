package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl string= "http://localhost:8000";

func main() {
	// getReq(baseUrl)
	// postReq(baseUrl)
	// postFormReq(baseUrl)
}

func getReq(baseUrl string) {
	response, err := http.Get(baseUrl+"/get");

	checkNilErr(err)

	fmt.Printf("type of response is %T\n", response) // *http.Response: reference of original response
	// fmt.Println(response)

	defer response.Body.Close() // important to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)

	// fmt.Println("databytes:", databytes)
	// fmt.Printf("data: %s", databytes)
	fmt.Println("content:", string(databytes))

	//OR

	var responseString strings.Builder
	bytecount, _ := responseString.Write(databytes)

	fmt.Println(bytecount)
	fmt.Println("from strings builder:", responseString.String())
}

func postReq(baseUrl string){
	endpoint := fmt.Sprintf("%s/post", baseUrl) //formats string

	respBody := strings.NewReader(`{
	"name": "muskan",
	"age": "22"
	}`)

	res, err := http.Post(endpoint, "application/json", respBody);

	checkNilErr(err)
	
	content, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	fmt.Println(string(content), res.StatusCode)
}

func postFormReq(baseUrl string){
	endpoint := fmt.Sprintf("%s/postform", baseUrl) //formats string

	respBody := url.Values{}

	respBody.Add("firstname", "Muskan")
	respBody.Add("lastname", "Shrivas")
	respBody.Add("post", "L7 (iykyk)")

	fmt.Println(respBody)
	fmt.Println(respBody.Encode())

	res, err := http.PostForm(endpoint, respBody);
	// postform auto handles encoding of form data

	checkNilErr(err)
	
	content, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	fmt.Println(string(content), res.StatusCode)
}

func checkNilErr (err error) {
	if err != nil {
		panic(err)
	}
}