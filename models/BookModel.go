package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Author Author             `json:"author"`
}
type Author struct {
	Name string `json:"name"`
}
