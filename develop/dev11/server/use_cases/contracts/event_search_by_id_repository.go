package contracts

import "d-alejandro/training-level2/develop/dev11/server/models"

type EventSearchByIDRepositoryContract interface {
	Make(id string) (*models.Event, error)
}
