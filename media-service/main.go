package main

import (
	"StorageService/internal/api"
	"StorageService/internal/db"
	_ "StorageService/internal/db"
	"StorageService/internal/handlers"
	"StorageService/internal/service"
	"log"
)

func main() {

	kafkaProducer, err := service.NewKafkaService()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandler(kafkaProducer)

	db.Init()
	serverInst := api.NewAPI(handler)
	httpServer := serverInst.RegisterRoutes()

	log.Fatal(httpServer.ListenAndServe())
}
