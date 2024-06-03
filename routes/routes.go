package routes

import (
    "context"
    "github.com/gin-gonic/gin"
    "go-kafka-mongo/config"
    "go-kafka-mongo/controllers"
    "go-kafka-mongo/repositories"
    "go-kafka-mongo/services"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Initialize database
    clientOptions := options.Client().ApplyURI(config.GetMongoURI())
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database("go_kafka_mongo")

    // Initialize repositories, services, and controllers
    bookRepo := repositories.NewBookRepository(db)
    bookService := services.NewBookService(bookRepo)
    bookController := controllers.NewBookController(bookService)

    // Set up routes
    router.POST("/books", bookController.CreateBook)
    router.GET("/books/:id", bookController.GetBook)
    router.PUT("/books/:id", bookController.UpdateBook)
    router.DELETE("/books/:id", bookController.DeleteBook)

    return router
}

