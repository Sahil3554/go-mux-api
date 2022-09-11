package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahil3554/go-mux-api/configs"
	"github.com/sahil3554/go-mux-api/middleware"
	"github.com/sahil3554/go-mux-api/routes"
)

func main() {
	configs.GetEnv()
	if configs.DB == nil {
		configs.ConnectDB()
	}
	router := mux.NewRouter()
	router.Use(middleware.CommonMiddleware)
	routes.Routes(router)
	fmt.Println("Starting Server at 8000 port ...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
