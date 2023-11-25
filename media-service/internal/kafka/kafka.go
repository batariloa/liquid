package kafka

import (
	"context"
	"encoding/json"
	"log"
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

func (ks *KafkaService) PublishUploadSongEvent(event UploadSongEvent) error {
	msgValue, err := json.Marshal(event)
	if err != nil {
		return err
	}

	log.Printf("Publishing event with value: %s", msgValue)

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
