package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayushhhh2999/mymodules/router"
)

func main() {
	fmt.Println("Starting the application...")
	r := router.Router()
	fmt.Println("Router initialized successfully")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running on port 4000")
}
