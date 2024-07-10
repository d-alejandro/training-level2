package contracts

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"time"
)

type EventCreationRepositoryContract interface {
	Make(name string, date time.Time) (*models.Event, error)
}
