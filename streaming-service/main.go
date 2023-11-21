package main

import (
	_ "github.com/batariloa/StreamingService/docs"
	"github.com/batariloa/StreamingService/internal/handler"
	"github.com/batariloa/StreamingService/internal/server"
	service "github.com/batariloa/StreamingService/internal/service/fetcher"
)

// @title Streaming Service API
func main() {

	fetcherService := service.NewFetcherService()
	streamHandler := handler.New(fetcherService)

	serverInstance := server.New(streamHandler)
	serverInstance.Start()
}
