package contracts

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

/*
EventCreationUseCaseContract contract
*/
type EventCreationUseCaseContract interface {
	Execute(*dto.EventRequestDTO) (*models.Event, error)
}
