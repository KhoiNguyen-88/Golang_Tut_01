package main

import (
    "go-kafka-mongo/config"
    "go-kafka-mongo/routes"
    "log"
)

func main() {
    config.InitConfig()
    router := routes.SetupRouter()
    err := router.Run(":8080")
    if err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}

