package use_cases

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/use_cases/contracts"
)

type EventDeletionUseCase struct {
	searchByIDRepository contracts.EventSearchByIDRepositoryContract
	deletionRepository   contracts.EventDeletionRepositoryContract
}

func NewEventDeletionUseCase(
	searchByIDRepository contracts.EventSearchByIDRepositoryContract,
	deletionRepository contracts.EventDeletionRepositoryContract,
) *EventDeletionUseCase {
	return &EventDeletionUseCase{
		searchByIDRepository: searchByIDRepository,
		deletionRepository:   deletionRepository,
	}
}

func (receiver *EventDeletionUseCase) Execute(id string) (*models.Event, error) {
	event, searchError := receiver.searchByIDRepository.Make(id)
	if searchError != nil {
		return nil, searchError
	}

	err := receiver.deletionRepository.Make(id)
	if err != nil {
		return nil, err
	}

	return event, nil
}
