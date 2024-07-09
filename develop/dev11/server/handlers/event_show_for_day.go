package handlers

import (
	"fmt"
	"net/http"
)

type EventShowForDayHandler struct {
}

func NewEventShowForDayHandler() *EventShowForDayHandler {
	return &EventShowForDayHandler{}
}

func (receiver *EventShowForDayHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "EventShowForDayHandler")
}
