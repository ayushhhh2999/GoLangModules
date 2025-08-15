package main

import (
	"fmt"
)

func main() {
	var veglist = []string{"carrot", "broccoli", "spinach"}
	fmt.Println("Vegetable List:", veglist)
	veglist = append(veglist, "cabbage", "cauliflower")
	fmt.Println("Updated Vegetable List:", veglist)
	veglist = append(veglist, "kale")
	fmt.Println("Vegetable List after addition:", veglist)
	veglist = append(veglist[:5])
	fmt.Println("Vegetable List after slicing:", veglist)
	veglist = append(veglist[0:2], veglist[3:]...)
	fmt.Println("Vegetable List after removal:", veglist)
	var fruits = make([]string, 5)
	fmt.Println("Fruits List:", fruits)
}
