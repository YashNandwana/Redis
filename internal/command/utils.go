package command

import "github.com/codecrafters-io/redis-starter-go/internal/storage"

func fetchAllKeys(db *storage.InMemoryDB) []string {
	var data []string
	for key, _ := range db.Strings {
		data = append(data, key)
	}
	return data
}
