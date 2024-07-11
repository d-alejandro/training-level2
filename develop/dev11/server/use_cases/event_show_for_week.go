package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
)

type EventShowForWeekUseCase struct {
	repository contracts.EventSearchForDayIntervalRepositoryContract
}

func NewEventShowForWeekUseCase(
	repository contracts.EventSearchForDayIntervalRepositoryContract,
) *EventShowForWeekUseCase {
	return &EventShowForWeekUseCase{repository}
}

func (receiver *EventShowForWeekUseCase) Execute(date string) ([]*models.Event, error) {
	parsedDate, parsedError := helpers.ParseDate(date)
	if parsedError != nil {
		return nil, parsedError
	}

	now, nowPackageError := helpers.GetNow(parsedDate)
	if nowPackageError != nil {
		return nil, nowPackageError
	}

	startDate := now.BeginningOfWeek()
	endDate := now.EndOfWeek()

	return receiver.repository.Make(&startDate, &endDate)
}
