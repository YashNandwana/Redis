package command

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type Config interface {
	FetchCommandAndServe(db *storage.InMemoryDB)
}

type config struct {
	Command string
	Action  string
}

func NewConfigCommandController(payload server.Request) *config {
	return &config{
		Command: payload.Args[1],
		Action:  payload.Args[2],
	}
}

func (c *config) FetchCommandAndServe(db *storage.InMemoryDB) (string, error) {
	var resp string
	switch c.Command {
	case "GET":
		resp = c.serveGetRequest(db)
	default:
		return resp, fmt.Errorf("unable to serve request")
	}
	return resp, nil
}

func (c *config) serveGetRequest(db *storage.InMemoryDB) string {
	switch c.Action {
	case "dir":
		return db.Args.ConfigDir
	default:
		return db.Args.ConfigDbFile
	}
}
