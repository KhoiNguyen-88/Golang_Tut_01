package controllers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go-kafka-mongo/models"
    "go-kafka-mongo/services"
    "net/http"
)

type BookController struct {
    Service *services.BookService
}

func NewBookController(service *services.BookService) *BookController {
    return &BookController{Service: service}
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.BindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := ctrl.Service.CreateBook(context.Background(), &book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

func (ctrl *BookController) GetBook(c *gin.Context) {
    id := c.Param("id")
    book, err := ctrl.Service.GetBook(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := c.BindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := ctrl.Service.UpdateBook(context.Background(), id, &book)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    err := ctrl.Service.DeleteBook(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

