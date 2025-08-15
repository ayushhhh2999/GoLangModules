package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "http://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("starting web server")
	request, err := http.Get(URL)
	errNilrequest(err)
	fmt.Println("response status code:", request.StatusCode)
	databytes, err := ioutil.ReadAll(request.Body)
	errNilrequest(err)
	fmt.Println("response body:", string(databytes))
	fmt.Printf("response type %T\n", request.Body)
	fmt.Println("response header:", request.Header)
	fmt.Println("response content length:", request.ContentLength)
	defer request.Body.Close()
}
func errNilrequest(err error) {
	if err != nil {
		panic(err)
	}
}
