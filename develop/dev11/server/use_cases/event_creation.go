package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/models"
)

type EventCreationUseCase struct {
}

func NewEventCreationUseCase() *EventCreationUseCase {
	return &EventCreationUseCase{}
}

func (receiver *EventCreationUseCase) Execute(dto *dto.EventRequestDTO) (*models.Event, error) {
	return nil, nil
}
