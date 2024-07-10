package repositories

import "d-alejandro/training-level2/develop/dev11/server/database"

type EventDeletionRepository struct {
	dbConnection database.CacheContract
}

func NewEventDeletionRepository(dbConnection database.CacheContract) *EventDeletionRepository {
	return &EventDeletionRepository{dbConnection}
}

func (receiver *EventDeletionRepository) Make(id string) error {
	return receiver.dbConnection.DeleteEvent(id)
}
