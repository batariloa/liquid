package service

import (
	"StorageService/internal/types"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var (
	producer *kafka.Writer
)

func Init() {
	brokerAddress := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")

	producer = &kafka.Writer{
		Addr:                   kafka.TCP(brokerAddress),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
}

func PublishUploadSongEvent(event types.UploadSongEvent) error {
	msgValue, err := json.Marshal(event)
	if err != nil {
		return err
	}

	log.Printf("Publishing event with value: %s", msgValue)

	err = producer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgValue,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	if producer != nil {
		return producer.Close()
	}
	return nil
}
