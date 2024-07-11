package handlers

import (
	"d-alejandro/training-level2/develop/dev11/server/handlers/contracts"
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"d-alejandro/training-level2/develop/dev11/server/validators"
	"net/http"
)

/*
EventCreationHandler structure
*/
type EventCreationHandler struct {
	useCase contracts.EventCreationUseCaseContract
}

/*
NewEventCreationHandler constructor
*/
func NewEventCreationHandler(useCase contracts.EventCreationUseCaseContract) *EventCreationHandler {
	return &EventCreationHandler{useCase}
}

/*
ServeHTTP method
*/
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

	eventPresenter := presenters.NewEventPresenter(responseWriter)
	eventPresenter.Present(event)
}
