package main

import (
	"StorageService/internal/db"
	_ "StorageService/internal/db"
	"StorageService/internal/server"
	"StorageService/internal/service"
	"log"
)

func main() {

	kafkaProducer, err := service.NewKafkaService()
	if err != nil {
		log.Fatal(err)
	}

	db.Init()
	serverInst := server.NewAPI(kafkaProducer)
	httpServer := serverInst.RegisterRoutes()

	log.Fatal(httpServer.ListenAndServe())
}
