package contracts

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"time"
)

type EventSearchForDayIntervalRepositoryContract interface {
	Make(startDate, endDate *time.Time) ([]*models.Event, error)
}
