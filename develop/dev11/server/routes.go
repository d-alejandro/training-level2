package server

import (
	"d-alejandro/training-level2/develop/dev11/server/bindings"
	"net/http"
)

func InitRoutes(serveMux *http.ServeMux) {
	handler := bindings.NewHandlerBinding()

	serveMux.Handle("POST /create_event", handler.EventCreationHandler)
	serveMux.Handle("POST /update_event/{id}", handler.EventUpdateHandler)
	serveMux.Handle("GET /events_for_day", handler.EventShowForDayHandler)
}