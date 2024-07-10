package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

type EventUpdateRepository struct {
	dbConnection database.CacheContract
}

func NewEventUpdateRepository(dbConnection database.CacheContract) *EventUpdateRepository {
	return &EventUpdateRepository{dbConnection}
}

func (receiver *EventUpdateRepository) Make(event *models.Event) error {
	return receiver.dbConnection.SetEvent(event.ID, event)
}
