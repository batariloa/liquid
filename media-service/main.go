package main

import (
	"StorageService/internal/db"
	_ "StorageService/internal/db"
	"StorageService/internal/server"
	"log"
)

func main() {

	db.Init()
	serverInst := server.NewServer()
	httpServer := serverInst.SetupServer()

	log.Fatal(httpServer.ListenAndServe())
}
