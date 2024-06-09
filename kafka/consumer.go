package kafka

import (
	"Golang_Tut_01/config"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.GetKafkaBroker()},
		Topic:   topic,
		GroupID: "Golang_Tut_01-group",
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
