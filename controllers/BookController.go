package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sahil3554/go-mux-api/configs"
	"github.com/sahil3554/go-mux-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
var userCollection *mongo.Collection = configs.GetCollection("books")

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	defer results.Close(ctx)
	var books []models.Book
	for results.Next(ctx) {
		var singleBook models.Book
		if err = results.Decode(&singleBook); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(map[string]string{"message": "internal server error"})
		}

		books = append(books, singleBook)
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(books)
}
func GetBook(res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book
	objId, converterr := primitive.ObjectIDFromHex(params["id"])

	defer cancel()
	if converterr != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Invalid Id",
		})
		return
	}
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&book)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"message": "No User Found!"})
		return
	}
	json.NewEncoder(res).Encode(book)
}
func CreateBook(res http.ResponseWriter, req *http.Request) {
	var requestBody models.Book
	json.NewDecoder(req.Body).Decode(&requestBody)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	newUser := models.Book{
		ID:     primitive.NewObjectID(),
		Name:   requestBody.Name,
		Author: requestBody.Author,
	}
	result, err := userCollection.InsertOne(ctx, newUser)

	defer cancel()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
	// if err != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
	// 	json.NewEncoder(rw).Encode(response)
	// 	return
	// }
	// Books = append(Books, Book{Name: requestBody.Name, ID: requestBody.ID, Author: &Author{Name: requestBody.Author.Name}})
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
