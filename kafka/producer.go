package kafka

import (
    "context"
    "github.com/segmentio/kafka-go"
    "go-kafka-mongo/config"
    "log"
)

func ProduceMessage(topic, message string) error {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{config.GetKafkaBroker()},
        Topic:   topic,
    })
    defer writer.Close()

    err := writer.WriteMessages(context.Background(), kafka.Message{
        Value: []byte(message),
    })
    if err != nil {
        log.Println("Failed to write message:", err)
        return err
    }
    return nil
}

