package bindings

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/repositories"
)

/*
RepositoryBinding structure
*/
type RepositoryBinding struct {
	EventCreationRepository             *repositories.EventCreationRepository
	EventSearchByIDRepository           *repositories.EventSearchByIDRepository
	EventUpdateRepository               *repositories.EventUpdateRepository
	EventDeletionRepository             *repositories.EventDeletionRepository
	EventSearchForDayRepository         *repositories.EventSearchForDayRepository
	EventSearchForDayIntervalRepository *repositories.EventSearchForDayIntervalRepository
}

/*
NewRepositoryBinding constructor
*/
func NewRepositoryBinding() *RepositoryBinding {
	dbConnection := database.GetDatabaseConnection()

	return &RepositoryBinding{
		EventCreationRepository:             repositories.NewEventCreationRepository(dbConnection),
		EventSearchByIDRepository:           repositories.NewEventSearchByIDRepository(dbConnection),
		EventUpdateRepository:               repositories.NewEventUpdateRepository(dbConnection),
		EventDeletionRepository:             repositories.NewEventDeletionRepository(dbConnection),
		EventSearchForDayRepository:         repositories.NewEventSearchForDayRepository(dbConnection),
		EventSearchForDayIntervalRepository: repositories.NewEventSearchForDayIntervalRepository(dbConnection),
	}
}
