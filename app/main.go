package main

import (
	"github.com/codecrafters-io/redis-starter-go/handler"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
	"log"
)

func main() {
	parser := server.NewRESPParser()
	router := handler.NewRouter()
	db := storage.NewInMemoryDB()

	// register the commands
	router.Register(db)

	srv := server.NewServer("0.0.0.0:6379", parser, router)
	if err := srv.Serve(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
