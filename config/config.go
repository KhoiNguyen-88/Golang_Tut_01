package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func InitConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetMongoURI() string {
    return os.Getenv("MONGO_URI")
}

func GetKafkaBroker() string {
    return os.Getenv("KAFKA_BROKER")
}

