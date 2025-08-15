package main

import (
	//imports for url
	"fmt"
	"net/url"
)

const URL = "http://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("starting web server")
	parts, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}
	//fmt.Println("scheme:", parts.Scheme)
	//fmt.Println("host:", parts.Host)
	//fmt.Println("path:", parts.Path)
	//fmt.Println("raw query:", parts.RawQuery)
	//fmt.Println("port:", parts.Port())
	//fmt.Println("query params:", parts.Query())
	params := parts.Query()
	for key, value := range params {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("Total number of query parameters:", len(params))
	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "go.dev",
		Path:     "docs",
		RawQuery: "search=golang",
	}
	fmt.Println("Parsed URL:", partsofurl)
	fmt.Println("URL string:", partsofurl.String())

}
