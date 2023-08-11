package main

import (
	"github.com/batariloa/StreamingService/internal/handler"
	"github.com/batariloa/StreamingService/internal/server"
	"github.com/batariloa/StreamingService/internal/service"
)

func main() {

	fetcherService := service.NewFetcherService()
	streamHandler := handler.New(fetcherService)

	serverInstance := server.New(streamHandler)
	serverInstance.Start()
}
