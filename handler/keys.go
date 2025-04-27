package handler

import (
	"github.com/codecrafters-io/redis-starter-go/internal/command"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type KeysHandler struct {
	db *storage.InMemoryDB
}

func NewKeysHandler(db *storage.InMemoryDB) *KeysHandler {
	return &KeysHandler{db: db}
}

func (h *KeysHandler) Handle(request server.Request) (server.Response, error) {
	key := command.NewKeysCommandController(request)
	resp, err := key.FetchPatternAndServe(h.db)
	if err != nil {
		return server.Response{}, err
	}
	return server.Response{Array: resp, IsArrayResponse: true}, nil
}
