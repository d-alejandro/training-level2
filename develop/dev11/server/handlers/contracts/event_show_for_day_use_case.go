package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

/*
EventShowForDayUseCaseContract contract
*/
type EventShowForDayUseCaseContract interface {
	Execute(date string) ([]*models.Event, error)
}
