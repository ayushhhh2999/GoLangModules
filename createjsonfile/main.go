package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course name"`
	Price    int
	Platform string `json:"website"`
	// Password field is omitted from JSON output
	// using `json:"-"` to ignore it
	Password string `json:"-"`
	// Tag field is optional and will be omitted if empty
	// using `omitempty` to skip it in JSON output
	Tag string `json:"tag,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in Golang")
	// Create a slice of courses
	courses := []course{
		{"Go", 299, "Udemy", "1234", "Programming"},
		{"Python", 199, "Coursera", "abcd", "Data Science"},
		{"JavaScript", 99, "Codecademy", "xyz", ""}, // Tag is empty, will be omitted
	}
	// Create JSON file
	createJSONFile(courses)
	fmt.Println("JSON file created successfully")
}
func createJSONFile(courses []course) {
	finalData, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(finalData))
}
