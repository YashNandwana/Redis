package handler

import (
	getCommand "github.com/codecrafters-io/redis-starter-go/internal/command/get"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type GetHandler struct {
	db *storage.InMemoryDB
}

func NewGetHandler(db *storage.InMemoryDB) *GetHandler {
	return &GetHandler{db: db}
}

func (h *GetHandler) Handle(req server.Request) (server.Response, error) {
	g := getCommand.NewGetCommandController(req.Args[1])
	value := g.GetKey(h.db)
	response := server.Response{Body: value}
	if value == "" {
		response.IsNull = true
	}
	return response, nil
}
