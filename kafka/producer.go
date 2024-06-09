package kafka

import (
	"Golang_Tut_01/config"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
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
