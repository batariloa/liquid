package main

import (
	"StorageService/internal/artist"
	"StorageService/internal/download"
	"StorageService/internal/handlers"
	"StorageService/internal/kafka"
	"StorageService/internal/server"
	"StorageService/internal/song"
	"StorageService/internal/upload"
	"StorageService/internal/util"
	"log"
)

func main() {
	dbCon := util.DbConnect()

	songDataRepo := song.NewPqlRepository(dbCon)
	artistRepository := artist.NewPqlRepository(dbCon)

	songDataService := song.NewSongDataService(songDataRepo)
	artistService := artist.NewArtistService(artistRepository)
	kafkaService := kafka.NewKafkaService()
	uploadService := upload.NewUploadService(kafkaService)
	downloadService := download.NewDownloadService(songDataService)

	storageHandler := handlers.NewSongHandler(*uploadService, *songDataService, *kafkaService, artistService)
	artistHandler := handlers.NewArtistHandler(artistService)
	downloadHandler := handlers.NewDownloadHandler(downloadService)

	serverInst := server.NewServer(storageHandler, artistHandler, downloadHandler)
	httpServer := serverInst.SetupServer()

	log.Fatal(httpServer.ListenAndServe())
}
