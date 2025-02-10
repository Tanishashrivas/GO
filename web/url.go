package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://example.com/learn:6060?query1=one&query2=two"

func main() {
	result, _ := url.Parse(myurl); // *url.URL

	fmt.Printf("Type of result %T\n", result)
	fmt.Println(result.Scheme, "\n", result.Path, "\n", result.RawQuery)
	fmt.Println(result.Port())

	queries := result.Query()

	for _, val := range queries {
		fmt.Println(val)
	}

	// constructing url again
	// make sure to use & with url for original reference
	url2 := &url.URL{
		Scheme: "https",
		Path: "demopath",
		Host: "myhost.com",
		RawQuery: "muskan=musk",
	}

	fmt.Println(url2.String())
}