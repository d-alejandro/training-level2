package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

/*
EventSearchByIDRepository structure
*/
type EventSearchByIDRepository struct {
	dbConnection database.CacheContract
}

/*
NewEventSearchByIDRepository constructor
*/
func NewEventSearchByIDRepository(dbConnection database.CacheContract) *EventSearchByIDRepository {
	return &EventSearchByIDRepository{dbConnection}
}

/*
Make method
*/
func (receiver *EventSearchByIDRepository) Make(id string) (*models.Event, error) {
	return receiver.dbConnection.GetEvent(id)
}
