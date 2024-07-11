package handlers

import (
	"d-alejandro/training-level2/develop/dev11/server/handlers/contracts"
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"d-alejandro/training-level2/develop/dev11/server/validators"
	"net/http"
)

type EventShowForWeekHandler struct {
	useCase contracts.EventShowForWeekUseCaseContract
}

func NewEventShowForWeekHandler(useCase contracts.EventShowForWeekUseCaseContract) *EventShowForWeekHandler {
	return &EventShowForWeekHandler{useCase}
}

func (receiver *EventShowForWeekHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	eventDateRequestValidator := validators.NewEventDateRequestValidator()
	errorPresenter := presenters.NewErrorPresenter(responseWriter)

	date, validationError := eventDateRequestValidator.Validate(request)
	if validationError != nil {
		errorPresenter.Present(http.StatusBadRequest, validationError)
		return
	}

	events, useCaseError := receiver.useCase.Execute(date)
	if useCaseError != nil {
		errorPresenter.Present(http.StatusServiceUnavailable, useCaseError)
		return
	}

	eventPresenter := presenters.NewEventListPresenter(responseWriter)
	eventPresenter.Present(events)
}
