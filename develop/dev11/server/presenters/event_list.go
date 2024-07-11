package presenters

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"encoding/json"
	"log"
	"net/http"
)

/*
EventListPresenter structure
*/
type EventListPresenter struct {
	responseWriter http.ResponseWriter
}

/*
NewEventListPresenter constructor
*/
func NewEventListPresenter(responseWriter http.ResponseWriter) *EventListPresenter {
	return &EventListPresenter{responseWriter}
}

/*
Present method
*/
func (presenter *EventListPresenter) Present(events []*models.Event) {
	presenter.responseWriter.Header().
		Set("Content-Type", "application/json")

	presenter.responseWriter.WriteHeader(http.StatusOK)

	response := struct {
		Result []*models.Event `json:"result"`
	}{
		Result: events,
	}

	encodedResponse, encodedError := json.Marshal(response)
	if encodedError != nil {
		log.Println(encodedError)
	}

	if _, err := presenter.responseWriter.Write(encodedResponse); err != nil {
		log.Println(err)
	}
}
