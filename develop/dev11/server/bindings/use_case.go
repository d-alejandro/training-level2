package bindings

import "d-alejandro/training-level2/develop/dev11/server/use_cases"

type UseCaseBinding struct {
	EventCreationUseCase *use_cases.EventCreationUseCase
}

func NewUseCaseBinding() *UseCaseBinding {
	repositoryBinding := NewRepositoryBinding()

	return &UseCaseBinding{
		EventCreationUseCase: use_cases.NewEventCreationUseCase(repositoryBinding.EventCreationRepository),
	}
}
