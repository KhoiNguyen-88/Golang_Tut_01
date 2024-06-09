package repositories

import (
	"Golang_Tut_01/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	Collection *mongo.Collection
}

func NewBookRepository(db *mongo.Database) *BookRepository {
	return &BookRepository{Collection: db.Collection("books")}
}

func (repo *BookRepository) CreateBook(ctx context.Context, book *models.Book) error {
	_, err := repo.Collection.InsertOne(ctx, book)
	return err
}

func (repo *BookRepository) GetBook(ctx context.Context, id string) (*models.Book, error) {
	var book models.Book
	err := repo.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	return &book, err
}

func (repo *BookRepository) UpdateBook(ctx context.Context, id string, book *models.Book) error {
	_, err := repo.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": book})
	return err
}

func (repo *BookRepository) DeleteBook(ctx context.Context, id string) error {
	_, err := repo.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
