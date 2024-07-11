package usecases

import (
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/usecases/contracts"
)

/*
EventShowForDayUseCase structure
*/
type EventShowForDayUseCase struct {
	repository contracts.EventSearchForDayRepositoryContract
}

/*
NewEventShowForDayUseCase constructor
*/
func NewEventShowForDayUseCase(repository contracts.EventSearchForDayRepositoryContract) *EventShowForDayUseCase {
	return &EventShowForDayUseCase{repository: repository}
}

/*
Execute method
*/
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
