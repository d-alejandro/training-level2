package usecases

import (
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/usecases/contracts"
)

/*
EventShowForMonthUseCase structure
*/
type EventShowForMonthUseCase struct {
	repository contracts.EventSearchForDayIntervalRepositoryContract
}

/*
NewEventShowForMonthUseCase constructor
*/
func NewEventShowForMonthUseCase(
	repository contracts.EventSearchForDayIntervalRepositoryContract,
) *EventShowForMonthUseCase {
	return &EventShowForMonthUseCase{repository}
}

/*
Execute method
*/
func (receiver *EventShowForMonthUseCase) Execute(date string) ([]*models.Event, error) {
	parsedDate, parsedError := helpers.ParseDate(date)
	if parsedError != nil {
		return nil, parsedError
	}

	now, nowPackageError := helpers.GetNow(parsedDate)
	if nowPackageError != nil {
		return nil, nowPackageError
	}

	startDate := now.BeginningOfMonth()
	endDate := now.EndOfMonth()

	return receiver.repository.Make(&startDate, &endDate)
}
