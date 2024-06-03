package models

type Book struct {
    ID          string `json:"id" bson:"_id,omitempty"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    Description string `json:"description"`
}

