package contracts

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

/*
EventUpdateUseCaseContract contract
*/
type EventUpdateUseCaseContract interface {
	Execute(id string, dto *dto.EventRequestDTO) (*models.Event, error)
}
