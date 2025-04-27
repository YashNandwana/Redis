// internal/storage/db.go
package storage

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/utils"
	"github.com/hdt3213/rdb/model"
	rdb "github.com/hdt3213/rdb/parser"

	"os"
	"path/filepath"
	"time"
)

type dbEntry[Type any] struct {
	Value     Type
	ExpiresAt time.Time
}

type InMemoryDB struct {
	Strings map[string]dbEntry[string]
	Hashes  map[string]dbEntry[map[string]string]
	Lists   map[string]dbEntry[[]string]
	Args    utils.CliArgs
}

func NewInMemoryDB(args utils.CliArgs) *InMemoryDB {
	return &InMemoryDB{
		Strings: make(map[string]dbEntry[string]),
		Hashes:  make(map[string]dbEntry[map[string]string]),
		Lists:   make(map[string]dbEntry[[]string]),
		Args:    args,
	}
}

func (db *InMemoryDB) Set(key, value string) error {
	db.Strings[key] = dbEntry[string]{
		Value:     value,
		ExpiresAt: time.Time{},
	}
	return nil
}

func (db *InMemoryDB) SetWithTTL(key, value string, ttl time.Duration) error {
	now := time.Now().Local()
	fmt.Println(now, ttl)
	db.Strings[key] = dbEntry[string]{
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

func (db *InMemoryDB) LoadRDB() {
	path := filepath.Join(db.Args.ConfigDir, db.Args.ConfigDbFile)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	dec := rdb.NewDecoder(f)
	dec.Parse(func(obj model.RedisObject) bool {
		key := obj.GetKey()
		var expiry time.Time

		if exp := obj.GetExpiration(); exp != nil {
			expiry = *exp
		}

		switch obj.GetType() {
		case rdb.StringType:
			o := obj.(*rdb.StringObject)

			if !expiry.IsZero() {
				ttl := time.Until(expiry)
				if ttl <= 0 {
					delete(db.Strings, key)
				} else {
					_ = db.SetWithTTL(key, string(o.Value), ttl)
				}
			} else {
				_ = db.Set(key, string(o.Value))
			}

		// TODO: handle other types (list, hash, etc.) in the same pattern:
		//   1. cast to *rdb.ListObject or *rdb.HashObject
		//   2. convert bytes → Go types
		//   3. if expiry.IsZero() { db.<Type>[key] = entry } else { compute ttl and call SetWithTTL... }

		default:
		}

		return true
	})
}
