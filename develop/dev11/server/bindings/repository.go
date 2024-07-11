package bindings

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/repositories"
)

type RepositoryBinding struct {
	EventCreationRepository   *repositories.EventCreationRepository
	EventSearchByIdRepository *repositories.EventSearchByIdRepository
	EventUpdateRepository     *repositories.EventUpdateRepository
	EventDeletionRepository   *repositories.EventDeletionRepository
	EventShowForDayRepository *repositories.EventShowForDayRepository
}

func NewRepositoryBinding() *RepositoryBinding {
	dbConnection := database.GetDatabaseConnection()

	return &RepositoryBinding{
		EventCreationRepository:   repositories.NewEventCreationRepository(dbConnection),
		EventSearchByIdRepository: repositories.NewEventSearchByIdRepository(dbConnection),
		EventUpdateRepository:     repositories.NewEventUpdateRepository(dbConnection),
		EventDeletionRepository:   repositories.NewEventDeletionRepository(dbConnection),
		EventShowForDayRepository: repositories.NewEventShowForDayRepository(dbConnection),
	}
}
