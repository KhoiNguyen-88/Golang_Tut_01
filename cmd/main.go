package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "Golang_Tut_01/config"
    "Golang_Tut_01/routes"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Access environment variables
    mongoHost := os.Getenv("MONGO_HOST")
    mongoPort := os.Getenv("MONGO_PORT")
    mongoUser := os.Getenv("MONGO_USER")
    mongoPass := os.Getenv("MONGO_PASS")
    mongoDB := os.Getenv("MONGO_DB")

    kafkaBroker := os.Getenv("KAFKA_BROKER")
    kafkaTopic := os.Getenv("KAFKA_TOPIC")

    appPort := os.Getenv("APP_PORT")
    logLevel := os.Getenv("LOG_LEVEL")

    // Print the variables (for debugging purposes)
    log.Println("MongoDB host:", mongoHost)
    log.Println("MongoDB port:", mongoPort)
    log.Println("MongoDB user:", mongoUser)
    log.Println("MongoDB password:", mongoPass)
    log.Println("MongoDB name:", mongoDB)

    log.Println("Kafka broker:", kafkaBroker)
    log.Println("Kafka topic:", kafkaTopic)

    log.Println("Application port:", appPort)
    log.Println("Log level:", logLevel)

    // Continue with the rest of your application logic

    config.InitConfig()
    router := routes.SetupRouter()
    err = router.Run(":8080")
    if err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
