package kafka

import (
    "context"
    "github.com/segmentio/kafka-go"
    "go-kafka-mongo/config"
    "log"
)

func ConsumeMessages(topic string) {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{config.GetKafkaBroker()},
        Topic:   topic,
        GroupID: "go-kafka-mongo-group",
    })
    defer reader.Close()

    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Println("Failed to read message:", err)
            continue
        }
        log.Printf("Received message: %s", string(msg.Value))
        // Process the message
    }
}

