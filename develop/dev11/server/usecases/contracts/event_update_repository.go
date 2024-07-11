package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

/*
EventUpdateRepositoryContract contract
*/
type EventUpdateRepositoryContract interface {
	Make(*models.Event) error
}
