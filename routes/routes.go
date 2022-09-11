package routes

import "github.com/gorilla/mux"

func Routes(router *mux.Router) {
	HomeRoutes(router)
	BookRoutes(router)
}
