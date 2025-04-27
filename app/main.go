package main

import (
	"flag"
	"log"

	"github.com/codecrafters-io/redis-starter-go/handler"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
	"github.com/codecrafters-io/redis-starter-go/utils"
)

func main() {
	// define flags
	var args utils.CliArgs
	flag.StringVar(&args.ConfigDir, "dir", "", "Directory to store RDB file")
	flag.StringVar(&args.ConfigDbFile, "dbfilename", "dump.rdb", "Name of the RDB file")
	flag.Parse()

	parser := server.NewRESPParser()
	router := handler.NewRouter()

	db := storage.NewInMemoryDB(args)
	db.LoadRDB()

	router.Register(db)

	srv := server.NewServer("0.0.0.0:6379", parser, router)
	if err := srv.Serve(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
