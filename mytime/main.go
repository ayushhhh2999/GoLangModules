package main

import (
	"fmt"
	"time"
)

func main() {
	currtime := time.Now()
	fmt.Println("Current time is:", currtime.Format("2006-01-02 15:04:05"))
	createdate := time.Date(2023, time.April, 12, 23, 0, 0, 0, time.UTC)
	fmt.Println("Created date is:", createdate.Format("2006-01-02 15:04:05"))
	fmt.Println("Time since created date:", currtime.Sub(createdate))
}
