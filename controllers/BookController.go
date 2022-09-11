package controllers

import (
	"encoding/json"
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

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(Books)
}
func GetBook(res http.ResponseWriter, req *http.Request) {
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
func CreateBook(res http.ResponseWriter, req *http.Request) {
	var requestBody Book
	json.NewDecoder(req.Body).Decode(&requestBody)
	requestBody.ID = strconv.Itoa(len(Books))
	Books = append(Books, Book{Name: requestBody.Name, ID: requestBody.ID, Author: &Author{Name: requestBody.Author.Name}})
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(requestBody)
}
func DeleteBook(res http.ResponseWriter, req *http.Request) {
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
func UpdateBook(res http.ResponseWriter, req *http.Request) {
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
