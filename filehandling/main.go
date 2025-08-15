package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "hello this is ayush singh"
	file, err := os.Create("./filehandling.txt")
	checkNilerror(err)
	length, err := io.WriteString(file, content)
	checkNilerror(err)
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile()
	data1, err := os.ReadFile("C:/Users/ayush/Desktop/demo.txt")
	fmt.Println("Data in file is:", string(data1))
}
func checkNilerror(err error) {
	if err != nil {
		panic(err)
	}
}
func readFile() {
	data, err := os.ReadFile("./filehandling.txt")
	checkNilerror(err)
	fmt.Println("Data in file is:", string(data))
}
