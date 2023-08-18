package consumer

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/batariloa/search-service/internal/model"
	"github.com/batariloa/search-service/internal/search"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumerService struct {
	reader        *kafka.Reader
	searchService *search.SearchService
}

func NewKafkaConsumerService(ss *search.SearchService) (*KafkaConsumerService, error) {

	brokerAddress := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")

	log.Printf("Kafka broker address: %s\n", brokerAddress)

	reader := kafka.NewReader(kafka.ReaderConfig{

		Brokers:  []string{brokerAddress},
		GroupID:  "search-consumer-group",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &KafkaConsumerService{
		reader:        reader,
		searchService: ss}, nil
}

func (c *KafkaConsumerService) ConsumeEvents() {
	defer c.reader.Close()

	for {
		msg, err := c.reader.FetchMessage(context.Background())
		if err != nil {
			log.Printf("Error fetching Kafka message: %v\n", err)
			continue
		}

		var song model.SongUploadEvent
		err = json.Unmarshal(msg.Value, &song)
		if err != nil {
			log.Println("Error decoding Kafka event:", err)
			continue
		}

		err = c.searchService.IndexSong(song)
		if err != nil {
			log.Println("Error indexing song:", err)
		}

		log.Printf("Consumed Kafka message: %s", song)
		c.reader.CommitMessages(context.Background(), msg)
	}
}
