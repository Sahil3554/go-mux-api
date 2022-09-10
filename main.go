package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}
type Author struct {
	Name string `json:"name"`
}

var Books []Book

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	fmt.Println("Starting Server at 8000 port ...")
	http.ListenAndServe(":8000", router)
}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(map[string]string{
		"message": "Working Successfully",
	})
}
func getAllBooks(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(Books)
}
func getBook(res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	for _, book := range Books {
		if book.ID == params["id"] {
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	json.NewEncoder(res).Encode(map[string]string{
		"message": "No Book Found with this id",
	})
}
func createBook(res http.ResponseWriter, req *http.Request) {
	var requestBody Book
	json.NewDecoder(req.Body).Decode(&requestBody)
	requestBody.ID = strconv.Itoa(len(Books))
	Books = append(Books, Book{Name: requestBody.Name, ID: requestBody.ID, Author: &Author{Name: requestBody.Author.Name}})
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(requestBody)
}
func deleteBook(res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	for index, book := range Books {
		if book.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	json.NewEncoder(res).Encode(map[string]string{
		"message": "No Book Found with this id",
	})
}
func updateBook(res http.ResponseWriter, req *http.Request) {
	var requestBody Book
	json.NewDecoder(req.Body).Decode(&requestBody)
	var params = mux.Vars(req)
	for index, book := range Books {
		if book.ID == params["id"] {
			Books[index].Name = requestBody.Name
			json.NewEncoder(res).Encode(map[string]string{
				"message": "Updated successfully!",
			})
			return
		}
	}
	json.NewEncoder(res).Encode(map[string]string{
		"message": "No Book Found with this id",
	})
}
