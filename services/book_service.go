package services

import (
	"Golang_Tut_01/models"
	"Golang_Tut_01/repositories"
	"context"
)

type BookService struct {
	Repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, book *models.Book) error {
	return s.Repo.CreateBook(ctx, book)
}

func (s *BookService) GetBook(ctx context.Context, id string) (*models.Book, error) {
	return s.Repo.GetBook(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, id string, book *models.Book) error {
	return s.Repo.UpdateBook(ctx, id, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	return s.Repo.DeleteBook(ctx, id)
}
