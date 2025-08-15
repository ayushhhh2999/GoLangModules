package main

import (
	"fmt"
)

func main() {
	type Person struct {
		Name     string
		Email    string
		Verified bool
		Age      int
	}
	var person1 Person = Person{"Ayush", "ayush.dev", true, 25}
	fmt.Println("Name:", person1.Name)
	fmt.Println("Email:", person1.Email)
	fmt.Println("Verified:", person1.Verified)
	fmt.Println("Age:", person1.Age)
	fmt.Printf("Person details: %+v\n", person1)
}
