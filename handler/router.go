package handler

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

// Router dispatches commands to registered Handlers
type Router struct {
	routes map[string]server.Handler
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]server.Handler)}
}

func (r *Router) Register(db *storage.InMemoryDB) {
	r.doRegister("PING", NewPingHandler())
	r.doRegister("ECHO", NewEchoHandler())
	r.doRegister("SET", NewSetHandler(db))
	r.doRegister("GET", NewGetHandler(db))
}

func (r *Router) doRegister(cmd string, h server.Handler) {
	r.routes[cmd] = h
}

func (r *Router) Handle(req server.Request) (server.Response, error) {
	h, ok := r.routes[req.Command]
	if !ok {
		return server.Response{}, fmt.Errorf("unknown command %q", req.Command)
	}
	return h.Handle(req)
}
