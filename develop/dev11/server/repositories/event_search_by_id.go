package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

type EventSearchByIdRepository struct {
	dbConnection database.CacheContract
}

func NewEventSearchByIdRepository(dbConnection database.CacheContract) *EventSearchByIdRepository {
	return &EventSearchByIdRepository{dbConnection}
}

func (receiver *EventSearchByIdRepository) Make(id string) (*models.Event, error) {
	return receiver.dbConnection.GetEvent(id)
}
