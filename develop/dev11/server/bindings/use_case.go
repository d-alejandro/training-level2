package bindings

import "d-alejandro/training-level2/develop/dev11/server/use_cases"

type UseCaseBinding struct {
	EventCreationUseCase   *use_cases.EventCreationUseCase
	EventUpdateUseCase     *use_cases.EventUpdateUseCase
	EventDeletionUseCase   *use_cases.EventDeletionUseCase
	EventShowForDayUseCase *use_cases.EventShowForDayUseCase
}

func NewUseCaseBinding() *UseCaseBinding {
	repositoryBinding := NewRepositoryBinding()

	return &UseCaseBinding{
		EventCreationUseCase: use_cases.NewEventCreationUseCase(repositoryBinding.EventCreationRepository),
		EventUpdateUseCase: use_cases.NewEventUpdateUseCase(
			repositoryBinding.EventSearchByIdRepository,
			repositoryBinding.EventUpdateRepository,
		),
		EventDeletionUseCase: use_cases.NewEventDeletionUseCase(
			repositoryBinding.EventSearchByIdRepository,
			repositoryBinding.EventDeletionRepository,
		),
		EventShowForDayUseCase: use_cases.NewEventShowForDayUseCase(repositoryBinding.EventShowForDayRepository),
	}
}
