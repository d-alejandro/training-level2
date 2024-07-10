package database

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"errors"
	"sync"
)

type Cache struct {
	sync.RWMutex
	data map[string]*models.Event
}

func NewCache() CacheContract {
	return &Cache{
		data: make(map[string]*models.Event),
	}
}

func (cache *Cache) SetEvent(id string, event *models.Event) error {
	cache.Lock()
	defer cache.Unlock()

	cache.data[id] = event

	return nil
}

func (cache *Cache) DeleteEvent(id string) error {
	cache.Lock()
	defer cache.Unlock()

	delete(cache.data, id)

	return nil
}

func (cache *Cache) LoadEvents() (map[string]*models.Event, error) {
	cache.RLock()
	defer cache.RUnlock()

	return cache.data, nil
}

func (cache *Cache) GetEvent(id string) (*models.Event, error) {
	cache.RLock()
	defer cache.RUnlock()

	event, found := cache.data[id]
	if !found {
		return nil, errors.New("event not found")
	}

	return event, nil
}
