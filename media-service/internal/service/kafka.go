package service

import (
	"StorageService/internal/types"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type EventPublisher interface {
	PublishUploadSongEvent(event types.UploadSongEvent) error
	Close() error
}

type KafkaService struct {
	producer *kafka.Writer
}

func NewKafkaService() (*KafkaService, error) {
	brokerAddress := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")

	producer := &kafka.Writer{
		Addr:                   kafka.TCP(brokerAddress),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	return &KafkaService{producer: producer}, nil
}

func (k *KafkaService) PublishUploadSongEvent(event types.UploadSongEvent) error {
	msgValue, err := json.Marshal(event)
	if err != nil {
		return err
	}

	log.Printf("Publishing event with value: %s", msgValue)

	err = k.producer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgValue,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaService) Close() error {
	if k.producer != nil {
		return k.producer.Close()
	}
	return nil
}
