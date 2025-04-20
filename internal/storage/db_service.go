// internal/storage/db.go
package storage

import (
	"fmt"
	"time"
)

type DbEntry[Type any] struct {
	Value     Type
	ExpiresAt time.Time
}

type InMemoryDB struct {
	Strings map[string]DbEntry[string]
	Hashes  map[string]DbEntry[map[string]string]
	Lists   map[string]DbEntry[[]string]
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Strings: make(map[string]DbEntry[string]),
		Hashes:  make(map[string]DbEntry[map[string]string]),
		Lists:   make(map[string]DbEntry[[]string]),
	}
}

func (db *InMemoryDB) Set(key, value string) error {
	db.Strings[key] = DbEntry[string]{
		Value:     value,
		ExpiresAt: time.Time{},
	}
	return nil
}

func (db *InMemoryDB) SetWithTTL(key, value string, ttl time.Duration) error {
	now := time.Now().Local()
	fmt.Println(now, ttl)
	db.Strings[key] = DbEntry[string]{
		Value:     value,
		ExpiresAt: now.Add(ttl),
	}
	return nil
}

func (db *InMemoryDB) Get(key string) (string, bool) {
	entry, ok := db.Strings[key]
	if !ok {
		return "", false
	}
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		delete(db.Strings, key)
		return "", false
	}
	return entry.Value, true
}
