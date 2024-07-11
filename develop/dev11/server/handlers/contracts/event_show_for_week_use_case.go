package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

/*
EventShowForWeekUseCaseContract contract
*/
type EventShowForWeekUseCaseContract interface {
	Execute(date string) ([]*models.Event, error)
}
