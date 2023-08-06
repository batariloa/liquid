package main

import (
	"StorageService/internal/handlers"
	"StorageService/internal/repository/artist"
	"StorageService/internal/repository/songdata"
	"StorageService/internal/server"
	"StorageService/internal/service"
	"StorageService/internal/util"
	"log"
)

func main() {
	dbCon := util.DbConnect()

	songDataRepo := songdata.NewPqlRepository(dbCon)
	artistRepository := artist.NewPqlRepository(dbCon)

	songDataService := service.NewSongDataService(songDataRepo)
	artistService := service.NewArtistService(artistRepository)
	kafkaService := service.NewKafkaService()
	uploadService := service.NewUploadService(kafkaService)
	downloadService := service.NewDownloadService(songDataService)

	storageHandler := handlers.NewSongHandler(uploadService, songDataService, kafkaService, artistService)
	artistHandler := handlers.NewArtistHandler(artistService)
	downloadHandler := handlers.NewDownloadHandler(downloadService)

	serverInst := server.NewServer(storageHandler, artistHandler, downloadHandler)
	httpServer := serverInst.SetupServer()

	log.Fatal(httpServer.ListenAndServe())
}
