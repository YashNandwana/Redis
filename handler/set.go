package handler

import (
	"github.com/codecrafters-io/redis-starter-go/internal/command"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type SetHandler struct {
	db *storage.InMemoryDB
}

func NewSetHandler(db *storage.InMemoryDB) *SetHandler {
	return &SetHandler{db: db}
}

func (h *SetHandler) Handle(req server.Request) (server.Response, error) {
	setKey := command.NewSetCommandController(req)
	err := setKey.SetKey(h.db)
	if err != nil {
		return server.Response{
			Body: "",
		}, err
	}
	return server.Response{Body: "OK"}, nil
}
