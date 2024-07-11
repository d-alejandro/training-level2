package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
)

type EventShowForDayUseCase struct {
	repository contracts.EventSearchForDayRepositoryContract
}

func NewEventShowForDayUseCase(repository contracts.EventSearchForDayRepositoryContract) *EventShowForDayUseCase {
	return &EventShowForDayUseCase{repository: repository}
}

func (receiver *EventShowForDayUseCase) Execute(date string) ([]*models.Event, error) {
	parsedDate, parsedError := helpers.ParseDate(date)
	if parsedError != nil {
		return nil, parsedError
	}

	events, err := receiver.repository.Make(parsedDate)
	if err != nil {
		return nil, err
	}

	return events, nil
}
