package handlers

import (
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"d-alejandro/training-level2/develop/dev11/server/validators"
	"fmt"
	"net/http"
)

type EventCreationHandler struct {
}

func NewEventCreationHandler() *EventCreationHandler {
	return &EventCreationHandler{}
}

func (receiver *EventCreationHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	eventRequestValidator := validators.NewEventRequestValidator()
	errorPresenter := presenters.NewErrorPresenter(responseWriter)

	if err := eventRequestValidator.Validate(request); err != nil {
		errorPresenter.Present(http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(responseWriter, "EventCreationHandler")
}
