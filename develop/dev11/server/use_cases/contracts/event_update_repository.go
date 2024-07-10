package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

type EventUpdateRepositoryContract interface {
	Make(*models.Event) error
}
