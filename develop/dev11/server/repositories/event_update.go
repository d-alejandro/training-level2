package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

/*
EventUpdateRepository structure
*/
type EventUpdateRepository struct {
	dbConnection database.CacheContract
}

/*
NewEventUpdateRepository constructor
*/
func NewEventUpdateRepository(dbConnection database.CacheContract) *EventUpdateRepository {
	return &EventUpdateRepository{dbConnection}
}

/*
Make method
*/
func (receiver *EventUpdateRepository) Make(event *models.Event) error {
	return receiver.dbConnection.SetEvent(event.ID, event)
}
