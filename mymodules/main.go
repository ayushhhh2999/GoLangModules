package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // gorilla/mux ek powerful HTTP router hai jo URL patterns handle karta hai
)

func main() {
	// Server start hone par ye message terminal par print hoga
	fmt.Println("Welcome to module in golang")

	// greeter function ko call kar rahe hain (ek normal greeting function hai)
	greeter()

	// ek new router bana rahe hain jo HTTP request ko handle karega
	r := mux.NewRouter()

	// jab GET request root path ("/") par aaye to serveHome function chale
	r.HandleFunc("/", serveHome).Methods("GET")

	// port 4000 par server ko start kar rahe hain
	// agar koi error aayi to program crash ho jayega aur error log hogi
	log.Fatal(http.ListenAndServe(":4000", r))
}

// ek simple function jo terminal par greeting message print karta hai
func greeter() {
	fmt.Println("Hello mod user, this is a greeter function")
}

// jab koi user "/" route ko GET method se hit kare to yeh function chalega
func serveHome(w http.ResponseWriter, r *http.Request) {
	// terminal par print karega ki home page hit hua
	fmt.Println("Home Page")

	// client ko HTML response bhej rahe hain
	w.Write([]byte("<h1>Welcome to my sexy golang page</h1><h2>Welcome to my sexy golang page</h2><h3>Welcome to my sexy golang page</h3>"))
}
