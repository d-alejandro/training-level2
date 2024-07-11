package bindings

import "d-alejandro/training-level2/develop/dev11/server/handlers"

/*
HandlerBinding structure
*/
type HandlerBinding struct {
	EventCreationHandler     *handlers.EventCreationHandler
	EventUpdateHandler       *handlers.EventUpdateHandler
	EventDeletionHandler     *handlers.EventDeletionHandler
	EventShowForDayHandler   *handlers.EventShowForDayHandler
	EventShowForWeekHandler  *handlers.EventShowForWeekHandler
	EventShowForMonthHandler *handlers.EventShowForMonthHandler
}

/*
NewHandlerBinding constructor
*/
func NewHandlerBinding() *HandlerBinding {
	useCaseBinding := NewUseCaseBinding()

	return &HandlerBinding{
		EventCreationHandler:     handlers.NewEventCreationHandler(useCaseBinding.EventCreationUseCase),
		EventUpdateHandler:       handlers.NewEventUpdateHandler(useCaseBinding.EventUpdateUseCase),
		EventDeletionHandler:     handlers.NewEventDeletionHandler(useCaseBinding.EventDeletionUseCase),
		EventShowForDayHandler:   handlers.NewEventShowForDayHandler(useCaseBinding.EventShowForDayUseCase),
		EventShowForWeekHandler:  handlers.NewEventShowForWeekHandler(useCaseBinding.EventShowForWeekUseCase),
		EventShowForMonthHandler: handlers.NewEventShowForMonthHandler(useCaseBinding.EventShowForMonthUseCase),
	}
}
