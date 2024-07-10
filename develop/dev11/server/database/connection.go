package database

import (
	"sync"
)

var (
	once     sync.Once
	database CacheContract
)

func GetDatabaseConnection() CacheContract {
	once.Do(func() {
		database = NewCache()
	})

	return database
}
