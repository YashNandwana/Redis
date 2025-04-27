package handler

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/internal/command"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type ConfigHandler struct {
	db *storage.InMemoryDB
}

func NewConfigHandler(db *storage.InMemoryDB) *ConfigHandler {
	return &ConfigHandler{
		db: db,
	}
}

func (h *ConfigHandler) Handle(request server.Request) (server.Response, error) {
	fmt.Println(request)
	conf := command.NewConfigCommandController(request)
	configData, err := conf.FetchCommandAndServe(h.db)
	if err != nil {
		return server.Response{}, err
	}
	var resp []string
	resp = append(resp, request.Args[2], configData)

	return server.Response{Array: resp, IsArrayResponse: true}, nil
}
