package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    int     `json:"courseId"`
	Coursename  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var DB []Course

func main() {
	fmt.Println("Welcome to creating API in Golang")
	r := mux.NewRouter()

	// Prepopulate fake DB
	DB = append(DB, Course{
		CourseId:    1,
		Coursename:  "Go Programming",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "John Doe",
			Website:  "https://johndoe.com",
		},
	})
	DB = append(DB, Course{
		CourseId:    2,
		Coursename:  "Python Programming",
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Jane Smith",
			Website:  "https://janesmith.com",
		},
	})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getcourses).Methods("GET")
	r.HandleFunc("/course/{id}", getcourse).Methods("GET")
	r.HandleFunc("/course", createcourse).Methods("POST")
	r.HandleFunc("/course/{id}", updatecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deletecourse).Methods("DELETE")

	fmt.Println("Server is running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func (c *Course) isempty() bool {
	return c.CourseId == 0 && c.Coursename == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Course API</h1>"))
}

func getcourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DB)
}

func getcourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}
	for _, course := range DB {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	http.Error(w, "Course not found", http.StatusNotFound)
}

func createcourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if course.isempty() {
		http.Error(w, "Course data is empty", http.StatusBadRequest)
		return
	}
	// naive: auto-increment ID (could improve)
	course.CourseId = len(DB) + 1
	DB = append(DB, course)
	json.NewEncoder(w).Encode(course)
}

func updatecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if course.isempty() {
		http.Error(w, "Course data is empty", http.StatusBadRequest)
		return
	}
	for i, c := range DB {
		if c.CourseId == id {
			course.CourseId = id // preserve id
			DB[i] = course
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	http.Error(w, "Course not found", http.StatusNotFound)
}

func deletecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}
	for i, course := range DB {
		if course.CourseId == id {
			DB = append(DB[:i], DB[i+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	http.Error(w, "Course not found", http.StatusNotFound)
}
