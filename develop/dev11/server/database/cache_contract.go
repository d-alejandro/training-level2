package database

import "d-alejandro/training-level2/develop/dev11/server/models"

type CacheContract interface {
	SetEvent(id string, event *models.Event) error
	DeleteEvent(id string) error
	LoadEvents() (map[string]*models.Event, error)
	GetEvent(id string) (*models.Event, error)
}
