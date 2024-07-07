package bindings

import "d-alejandro/training-level2/develop/dev11/server/handlers"

type HandlerBinding struct {
	EventCreationHandler   *handlers.EventCreationHandler
	EventUpdateHandler     *handlers.EventUpdateHandler
	EventShowForDayHandler *handlers.EventShowForDayHandler
}

func NewHandlerBinding() *HandlerBinding {
	return &HandlerBinding{
		EventCreationHandler:   handlers.NewEventCreationHandler(),
		EventUpdateHandler:     handlers.NewEventUpdateHandler(),
		EventShowForDayHandler: handlers.NewEventShowForDayHandler(),
	}
}
