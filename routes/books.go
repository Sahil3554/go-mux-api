package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahil3554/go-mux-api/controllers"
)

func BookRoutes(router *mux.Router) {
	//All routes related to users comes here
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
}
