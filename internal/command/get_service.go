package command

import (
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
)

type Get interface {
	GetString() string
}

type getKey struct {
	Key string
}

func NewGetCommandController(key string) *getKey {
	g := &getKey{
		Key: key,
	}
	return g
}

func (g *getKey) GetKey(db *storage.InMemoryDB) string {
	value, _ := db.Get(g.Key)
	return value
}
