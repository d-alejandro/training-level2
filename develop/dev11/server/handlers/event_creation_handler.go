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
	errorPresenter := presenters.NewErrorPresenter(responseWriter)

	if err := request.ParseForm(); err != nil {
		errorPresenter.Present(http.StatusBadRequest, err)
		return
	}

	eventRequestValidator := validators.NewEventRequestValidator(request)
	if err := eventRequestValidator.Validate(); err != nil {
		errorPresenter.Present(http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(responseWriter, "EventCreationHandler")
}
