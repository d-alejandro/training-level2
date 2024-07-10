package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
	"time"
)

type EventUpdateUseCase struct {
	searchByIDRepository contracts.EventSearchByIDRepositoryContract
	updateRepository     contracts.EventUpdateRepositoryContract
}

func NewEventUpdateUseCase(
	searchByIDRepository contracts.EventSearchByIDRepositoryContract,
	updateRepository contracts.EventUpdateRepositoryContract,
) *EventUpdateUseCase {
	return &EventUpdateUseCase{
		searchByIDRepository: searchByIDRepository,
		updateRepository:     updateRepository,
	}
}

func (receiver *EventUpdateUseCase) Execute(id string, dto *dto.EventRequestDTO) (*models.Event, error) {
	event, searchError := receiver.searchByIDRepository.Make(id)
	if searchError != nil {
		return nil, searchError
	}

	const layout = "2006-01-02"

	date, parseError := time.Parse(layout, dto.GetDate())
	if parseError != nil {
		return nil, parseError
	}

	event.Name = dto.GetName()
	event.Date = date

	err := receiver.updateRepository.Make(event)

	if err != nil {
		return nil, err
	}

	return event, nil
}
