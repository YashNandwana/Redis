package handler

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Handle(req server.Request) (server.Response, error) {
	if len(req.Args) < 1 {
		return server.Response{}, fmt.Errorf("no argument to echo")
	}
	return server.Response{Body: req.Args[len(req.Args)-1]}, nil
}
