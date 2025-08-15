package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 👋 Main function — program yahin se start hota hai
func main() {
	fmt.Println("Client program started...")

	// ✅ Uncomment one by one and test each request type
	// sendGetRequest()
	// sendPostJSONRequest()
	sendPostFormRequest()
}

/*
🔷 GET REQUEST — Kya hoti hai?
--------------------------------
- Jab hum kisi server se sirf data "mangte" hain, bina kuch bheje.
- Browser mein jab tum URL type karte ho, wo GET request hi hoti hai.
- Query parameters ke zariye data bheja ja sakta hai (like ?name=ayush)
*/
func sendGetRequest() {
	url := "http://localhost:8000"

	// Server ko GET request bhej rahe hain
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // response complete hone ke baad close karna zaroori hai

	// Response body padhein
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ GET Request Response:")
	fmt.Println("Status Code:", response.Status)
	fmt.Println("Content Length:", response.ContentLength)
	fmt.Println("Response Body:", string(body))
}

/*
🔷 POST REQUEST with JSON body
----------------------------------
- Jab hum server ko data bhejna chahte hain (like user input, login data, etc)
- Data "body" mein bhejte hain JSON format mein
- Mostly APIs mein POST ka use hota hai jab kuch insert, update ya login karna hota hai
*/
func sendPostJSONRequest() {
	url := "http://localhost:8000"

	// JSON data jo body mein bhejna hai
	jsonPayload := strings.NewReader(`
		{
			"coursename": "lets go with golang",
			"price": "0",
			"platform": "learncodeonline"
		}
	`)

	// POST request bhejna with content-type = application/json
	response, err := http.Post(url, "application/json", jsonPayload)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Response read karein
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ POST JSON Request Response:")
	fmt.Println("Status Code:", response.Status)
	fmt.Println("Response Body:", string(body))
}

/*
🔷 POST FORM REQUEST — Jaise HTML form submit hota hai
---------------------------------------------------------
- Jab hum HTML form bhar ke submit karte hain, wo backend tak POST Form ke zariye jaata hai
- Data key-value pairs ke form mein hota hai (x-www-form-urlencoded)
*/
func sendPostFormRequest() {
	url := "http://localhost:8000/postform"

	// Form ke data ko prepare karna (key-value form mein)
	formData := url.Values{}
	formData.Add("name", "ayush")
	formData.Add("age", "22")
	formData.Add("language", "Go")

	// Form data ke saath POST request bhejna
	response, err := http.PostForm(url, formData)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Server se response read karna
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ POST Form Request Response:")
	fmt.Println("Status Code:", response.Status)
	fmt.Println("Response Body:", string(body))
}
