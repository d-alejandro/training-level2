package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

type EventShowForMonthUseCaseContract interface {
	Execute(date string) ([]*models.Event, error)
}
