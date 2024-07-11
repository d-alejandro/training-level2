package bindings

import "d-alejandro/training-level2/develop/dev11/server/usecases"

/*
UseCaseBinding structure
*/
type UseCaseBinding struct {
	EventCreationUseCase     *usecases.EventCreationUseCase
	EventUpdateUseCase       *usecases.EventUpdateUseCase
	EventDeletionUseCase     *usecases.EventDeletionUseCase
	EventShowForDayUseCase   *usecases.EventShowForDayUseCase
	EventShowForWeekUseCase  *usecases.EventShowForWeekUseCase
	EventShowForMonthUseCase *usecases.EventShowForMonthUseCase
}

/*
NewUseCaseBinding constructor
*/
func NewUseCaseBinding() *UseCaseBinding {
	repositoryBinding := NewRepositoryBinding()

	return &UseCaseBinding{
		EventCreationUseCase: usecases.NewEventCreationUseCase(repositoryBinding.EventCreationRepository),
		EventUpdateUseCase: usecases.NewEventUpdateUseCase(
			repositoryBinding.EventSearchByIDRepository,
			repositoryBinding.EventUpdateRepository,
		),
		EventDeletionUseCase: usecases.NewEventDeletionUseCase(
			repositoryBinding.EventSearchByIDRepository,
			repositoryBinding.EventDeletionRepository,
		),
		EventShowForDayUseCase: usecases.NewEventShowForDayUseCase(repositoryBinding.EventSearchForDayRepository),
		EventShowForWeekUseCase: usecases.NewEventShowForWeekUseCase(
			repositoryBinding.EventSearchForDayIntervalRepository,
		),
		EventShowForMonthUseCase: usecases.NewEventShowForMonthUseCase(
			repositoryBinding.EventSearchForDayIntervalRepository,
		),
	}
}
