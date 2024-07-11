package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

/*
EventSearchByIDRepositoryContract contract
*/
type EventSearchByIDRepositoryContract interface {
	Make(id string) (*models.Event, error)
}
