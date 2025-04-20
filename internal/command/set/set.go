package set

import (
	"github.com/codecrafters-io/redis-starter-go/internal/server"
	"github.com/codecrafters-io/redis-starter-go/internal/storage"
	"strconv"
	"strings"
	"time"
)

type Set interface {
	Update() string
}

type setKey struct {
	Key   string
	Value string
	EX    time.Duration
	PX    time.Duration
}

func NewSetCommandController(payload server.Request) *setKey {
	var s setKey
	switch len(payload.Args) {
	case 3:
		s = setKey{
			Key:   payload.Args[1],
			Value: payload.Args[2],
		}
	default:
		s = setKey{
			Key:   payload.Args[1],
			Value: payload.Args[2],
		}
		ttl, _ := strconv.ParseFloat(strings.TrimSpace(payload.Args[4]), 64)
		if payload.Args[3] == "px" {
			s.PX = time.Duration(ttl) * time.Millisecond
		} else {
			s.EX = time.Duration(ttl) * time.Millisecond * 1000
		}
	}
	return &s
}

func (s *setKey) SetKey(db *storage.InMemoryDB) error {
	if s.PX == 0 && s.EX == 0 {
		err := db.Set(s.Key, s.Value)
		if err != nil {
			return err
		}
		return nil
	} else {
		if s.PX != 0 {
			err := db.SetWithTTL(s.Key, s.Value, s.PX)
			if err != nil {
				return err
			}
		} else {
			err := db.SetWithTTL(s.Key, s.Value, s.EX)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
