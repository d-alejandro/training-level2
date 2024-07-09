package handlers

import (
	"fmt"
	"net/http"
)

type EventUpdateHandler struct {
}

func NewEventUpdateHandler() *EventUpdateHandler {
	return &EventUpdateHandler{}
}

func (receiver *EventUpdateHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "EventUpdateHandler")
}
