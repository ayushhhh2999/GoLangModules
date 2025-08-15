//package main

// func main() {
// fmt.Println("hello nigga")
// main1()
// vars()
// cons()}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter the rating")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")
	rating, _ := reader.ReadString('\n')
	fmt.Println("You entered:", rating)

}
