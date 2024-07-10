package contracts

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
)

type EventDeletionUseCaseContract interface {
	Execute(id string) (*models.Event, error)
}
