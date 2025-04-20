package handler

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Handle(req server.Request) (server.Response, error) {
	fmt.Println(req)
	return server.Response{Body: "PONG"}, nil
}
