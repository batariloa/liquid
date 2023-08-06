package service

import (
	"StorageService/internal/datastruct"
	"context"
	"encoding/json"
	"os"

	"github.com/segmentio/kafka-go"
)

type KafkaService struct {
	Producer *kafka.Writer
}

func NewKafkaService() *KafkaService {
	brokerAddress := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")

	return &KafkaService{
		Producer: &kafka.Writer{
			Addr:                   kafka.TCP(brokerAddress),
			Topic:                  topic,
			AllowAutoTopicCreation: true,
		},
	}
}

func (ks *KafkaService) PublishEvent(data datastruct.UploadKafkaEvent) error {
	msgValue, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ks.Producer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgValue,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (ks *KafkaService) Close() error {
	if ks.Producer != nil {
		return ks.Producer.Close()
	}
	return nil
}
