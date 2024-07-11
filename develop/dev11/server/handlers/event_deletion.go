package handlers

import (
	"d-alejandro/training-level2/develop/dev11/server/handlers/contracts"
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"net/http"
)

/*
EventDeletionHandler structure
*/
type EventDeletionHandler struct {
	useCase contracts.EventDeletionUseCaseContract
}

/*
NewEventDeletionHandler constructor
*/
func NewEventDeletionHandler(useCase contracts.EventDeletionUseCaseContract) *EventDeletionHandler {
	return &EventDeletionHandler{useCase}
}

/*
ServeHTTP method
*/
func (receiver *EventDeletionHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	event, useCaseError := receiver.useCase.Execute(id)
	if useCaseError != nil {
		errorPresenter := presenters.NewErrorPresenter(responseWriter)
		errorPresenter.Present(http.StatusServiceUnavailable, useCaseError)
		return
	}

	eventPresenter := presenters.NewEventPresenter(responseWriter)
	eventPresenter.Present(event)
}
