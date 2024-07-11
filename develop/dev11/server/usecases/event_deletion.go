package usecases

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"d-alejandro/training-level2/develop/dev11/server/usecases/contracts"
)

/*
EventDeletionUseCase structure
*/
type EventDeletionUseCase struct {
	searchByIDRepository contracts.EventSearchByIDRepositoryContract
	deletionRepository   contracts.EventDeletionRepositoryContract
}

/*
NewEventDeletionUseCase constructor
*/
func NewEventDeletionUseCase(
	searchByIDRepository contracts.EventSearchByIDRepositoryContract,
	deletionRepository contracts.EventDeletionRepositoryContract,
) *EventDeletionUseCase {
	return &EventDeletionUseCase{
		searchByIDRepository: searchByIDRepository,
		deletionRepository:   deletionRepository,
	}
}

/*
Execute method
*/
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
