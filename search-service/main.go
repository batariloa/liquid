package main

import (
	"log"

	"github.com/batariloa/search-service/internal/consumer"
	"github.com/batariloa/search-service/internal/handler"
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
		log.Println("Cannot establish a Kafka connection.")
	}
	go kafkaConsumer.ConsumeEvents()

	searchHandler := handler.NewSearchHandler(searchService)
	serverInstance := server.NewServer(searchHandler)
	serverInstance.Run(":8085")
}
