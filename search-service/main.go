package main

import (
	"fmt"

	"github.com/batariloa/search-service/internal/consumer"
	"github.com/batariloa/search-service/internal/search"
	"github.com/batariloa/search-service/internal/server"
)

func main() {

	index, err := search.InitializeBleveIndex()
	if err != nil {
		panic(err)
	}

	searchService := search.NewSearchService(index)

	kafkaConsumer, err := consumer.NewKafkaConsumerService(searchService)
	if err != nil {
		fmt.Println("Cannot establish a Kafka connection.")
	}

	kafkaConsumer.ConsumeEvents()

	serverInstance := server.NewServer(searchService)
	serverInstance.Run("localhost:8085")
}
