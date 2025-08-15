package main

import (
	"fmt"
	"math"
)

func cons() {
	const n = 6000000000
	const d = 3e12 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
