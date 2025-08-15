package main

import (
	"fmt"
)

func main() {
	var number1 *int
	fmt.Println(number1)
	number2 := 10
	var ptr = &number2
	fmt.Println(ptr)
	*ptr = *ptr * 2
	fmt.Println("Value of number2 after dereferencing and modifying:", number2)
}
