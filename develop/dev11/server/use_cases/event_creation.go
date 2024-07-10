package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
)

type EventCreationUseCase struct {
	repository contracts.EventCreationRepositoryContract
}

func NewEventCreationUseCase(repository contracts.EventCreationRepositoryContract) *EventCreationUseCase {
	return &EventCreationUseCase{repository}
}

func (receiver *EventCreationUseCase) Execute(dto *dto.EventRequestDTO) (*models.Event, error) {
	date, err := helpers.ParseDate(dto.GetDate())

	if err != nil {
		return nil, err
	}

	return receiver.repository.Make(dto.GetName(), date)
}
