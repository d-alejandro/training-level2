package handlers

import (
	"d-alejandro/training-level2/develop/dev11/server/handlers/contracts"
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"d-alejandro/training-level2/develop/dev11/server/validators"
	"net/http"
)

/*
EventUpdateHandler structure
*/
type EventUpdateHandler struct {
	useCase contracts.EventUpdateUseCaseContract
}

/*
NewEventUpdateHandler constructor
*/
func NewEventUpdateHandler(useCase contracts.EventUpdateUseCaseContract) *EventUpdateHandler {
	return &EventUpdateHandler{useCase}
}

/*
ServeHTTP method
*/
func (receiver *EventUpdateHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	eventRequestValidator := validators.NewEventRequestValidator()
	errorPresenter := presenters.NewErrorPresenter(responseWriter)

	eventRequestDTO, validationError := eventRequestValidator.Validate(request)
	if validationError != nil {
		errorPresenter.Present(http.StatusBadRequest, validationError)
		return
	}

	id := request.PathValue("id")

	event, useCaseError := receiver.useCase.Execute(id, eventRequestDTO)
	if useCaseError != nil {
		errorPresenter.Present(http.StatusServiceUnavailable, useCaseError)
		return
	}

	eventPresenter := presenters.NewEventPresenter(responseWriter)
	eventPresenter.Present(event)
}
