package usecases

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/helpers"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/usecases/contracts"
)

/*
EventUpdateUseCase structure
*/
type EventUpdateUseCase struct {
	searchByIDRepository contracts.EventSearchByIDRepositoryContract
	updateRepository     contracts.EventUpdateRepositoryContract
}

/*
NewEventUpdateUseCase constructor
*/
func NewEventUpdateUseCase(
	searchByIDRepository contracts.EventSearchByIDRepositoryContract,
	updateRepository contracts.EventUpdateRepositoryContract,
) *EventUpdateUseCase {
	return &EventUpdateUseCase{
		searchByIDRepository: searchByIDRepository,
		updateRepository:     updateRepository,
	}
}

/*
Execute method
*/
func (receiver *EventUpdateUseCase) Execute(id string, dto *dto.EventRequestDTO) (*models.Event, error) {
	event, searchError := receiver.searchByIDRepository.Make(id)
	if searchError != nil {
		return nil, searchError
	}

	date, parseError := helpers.ParseDate(dto.GetDate())
	if parseError != nil {
		return nil, parseError
	}

	event.Name = dto.GetName()
	event.Date = *date

	err := receiver.updateRepository.Make(event)

	if err != nil {
		return nil, err
	}

	return event, nil
}
