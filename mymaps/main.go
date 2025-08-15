package main

import (
	"fmt"
)

func main() {
	hashmap := make(map[int]string)
	hashmap[1] = "go"
	hashmap[2] = "python"
	hashmap[3] = "java"
	hashmap[4] = "c++"
	hashmap[5] = "javascript"
	hashmap[6] = "rust"
	for key, value := range hashmap {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("Total number of languages:", len(hashmap))
}
