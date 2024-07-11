package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
)

type EventShowForMonthUseCase struct {
	repository contracts.EventSearchForDayIntervalRepositoryContract
}

func NewEventShowForMonthUseCase(
	repository contracts.EventSearchForDayIntervalRepositoryContract,
) *EventShowForMonthUseCase {
	return &EventShowForMonthUseCase{repository}
}

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
