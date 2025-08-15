package router

import (
	"github.com/ayushhhh2999/mymodules/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/movies", controller.Getallthemovies).Methods("GET")
	r.HandleFunc("/movies", controller.Addthemovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.Getonethemovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.Markthewatched).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.Deletetheone).Methods("DELETE")
	r.HandleFunc("/movies", controller.Deleteallthemovies).Methods("DELETE")

	return r
}
