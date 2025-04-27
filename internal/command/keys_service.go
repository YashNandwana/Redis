package command

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type Keys interface {
	FetchPatternAndServe(db *storage.InMemoryDB) (string, error)
}

type keys struct {
	Pattern string
}

func NewKeysCommandController(payload server.Request) *keys {
	return &keys{
		Pattern: payload.Args[1],
	}
}

func (k *keys) FetchPatternAndServe(db *storage.InMemoryDB) ([]string, error) {
	var resp []string
	switch k.Pattern {
	case "*":
		resp = fetchAllKeys(db)
	default:
		return resp, fmt.Errorf("unable to serve request")
	}
	return resp, nil
}
