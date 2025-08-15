package main

import "fmt"

func vars() {
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	var d = true
	var e int
	f := 3
	fmt.Println(a, b, c, d, e, f)
}
