package repositories

import "d-alejandro/training-level2/develop/dev11/server/database"

/*
EventDeletionRepository structure
*/
type EventDeletionRepository struct {
	dbConnection database.CacheContract
}

/*
NewEventDeletionRepository constructor
*/
func NewEventDeletionRepository(dbConnection database.CacheContract) *EventDeletionRepository {
	return &EventDeletionRepository{dbConnection}
}

/*
Make method
*/
func (receiver *EventDeletionRepository) Make(id string) error {
	return receiver.dbConnection.DeleteEvent(id)
}
