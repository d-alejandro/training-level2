package handlers

import (
	"d-alejandro/training-level2/develop/dev11/server/handlers/contracts"
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"d-alejandro/training-level2/develop/dev11/server/validators"
	"fmt"
	"net/http"
)

type EventCreationHandler struct {
	useCase contracts.EventCreationUseCaseContract
}

func NewEventCreationHandler(useCase contracts.EventCreationUseCaseContract) *EventCreationHandler {
	return &EventCreationHandler{useCase}
}

func (receiver *EventCreationHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	eventRequestValidator := validators.NewEventRequestValidator()
	errorPresenter := presenters.NewErrorPresenter(responseWriter)

	eventRequestDTO, validationError := eventRequestValidator.Validate(request)
	if validationError != nil {
		errorPresenter.Present(http.StatusBadRequest, validationError)
		return
	}

	event, useCaseError := receiver.useCase.Execute(eventRequestDTO)
	if useCaseError != nil {
		errorPresenter.Present(http.StatusServiceUnavailable, useCaseError)
		return
	}

	_ = event

	fmt.Fprint(responseWriter, "EventCreationHandler")
}
