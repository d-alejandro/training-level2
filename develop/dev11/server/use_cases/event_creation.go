package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
	"time"
)

type EventCreationUseCase struct {
	repository contracts.EventCreationRepositoryContract
}

func NewEventCreationUseCase(repository contracts.EventCreationRepositoryContract) *EventCreationUseCase {
	return &EventCreationUseCase{repository}
}

func (receiver *EventCreationUseCase) Execute(dto *dto.EventRequestDTO) (*models.Event, error) {
	const layout = "2006-01-02"

	date, err := time.Parse(layout, dto.GetDate())

	if err != nil {
		return nil, err
	}

	return receiver.repository.Make(dto.GetName(), date)
}
