package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the rating:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	rating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	rating += 1

	fmt.Println("Incremented rating:", rating)
	fmt.Println("Thank you for your input!")
}
