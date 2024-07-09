package bindings

import "d-alejandro/training-level2/develop/dev11/server/use_cases"

type UseCaseBinding struct {
	EventCreationUseCase *use_cases.EventCreationUseCase
}

func NewUseCaseBinding() *UseCaseBinding {
	return &UseCaseBinding{
		EventCreationUseCase: use_cases.NewEventCreationUseCase(),
	}
}
