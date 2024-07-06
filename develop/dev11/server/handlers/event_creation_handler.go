package handlers

import (
	"fmt"
	"net/http"
)

type EventCreationHandler struct {
}

func NewEventCreationHandler() *EventCreationHandler {
	return &EventCreationHandler{}
}

func (receiver *EventCreationHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "EventCreationHandler")
}
