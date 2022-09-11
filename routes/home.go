package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahil3554/go-mux-api/controllers"
)

func HomeRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
}
