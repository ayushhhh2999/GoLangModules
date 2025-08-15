package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"course name"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tag      string `json:"tag,omitempty"`
}

func main() {
	fmt.Println("Welcome to consuming JSON data in Golang")

	// JSON data (as if from web)
	jsonData := []byte(`
		[
			{
				"course name": "Go",
				"Price": 299,
				"website": "Udemy",
				"tag": "Programming"
			},
			{
				"course name": "Python",
				"Price": 199,
				"website": "Coursera",
				"tag": "Data Science"
			},
			{
				"course name": "JavaScript",
				"Price": 99,
				"website": "Codecademy",
				"tag": "Web Development"
			}
		]
	`)

	// Function call
	decodingjson(jsonData)
}

func decodingjson(data_from_web []byte) {
	check := json.Valid(data_from_web)
	if check {
		var courses []Course
		err := json.Unmarshal(data_from_web, &courses)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		fmt.Printf("Decoded struct slice: %+v\n", courses)
	} else {
		fmt.Println("JSON was not valid")
	}

	// Decoding JSON into a map (just showing first object)
	fmt.Println("Decoding first object into a map")
	var newcourse []map[string]interface{}
	err := json.Unmarshal(data_from_web, &newcourse)
	if err != nil {
		fmt.Println("Error unmarshalling to map:", err)
		return
	}
	fmt.Println("First course name is:", newcourse[0]["course name"])
}
