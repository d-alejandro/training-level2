package contracts

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"time"
)

/*
EventSearchForDayRepositoryContract contract
*/
type EventSearchForDayRepositoryContract interface {
	Make(date *time.Time) ([]*models.Event, error)
}
