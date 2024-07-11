package usecases

import (
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/usecases/contracts"
)

/*
EventShowForWeekUseCase structure
*/
type EventShowForWeekUseCase struct {
	repository contracts.EventSearchForDayIntervalRepositoryContract
}

/*
NewEventShowForWeekUseCase constructor
*/
func NewEventShowForWeekUseCase(
	repository contracts.EventSearchForDayIntervalRepositoryContract,
) *EventShowForWeekUseCase {
	return &EventShowForWeekUseCase{repository}
}

/*
Execute method
*/
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
