package presenters

import (
	"d-alejandro/training-level2/develop/dev11/server/models"
	"encoding/json"
	"log"
	"net/http"
)

type EventPresenter struct {
	responseWriter http.ResponseWriter
}

func NewEventPresenter(responseWriter http.ResponseWriter) *EventPresenter {
	return &EventPresenter{responseWriter}
}

func (presenter *EventPresenter) Present(event *models.Event) {
	presenter.responseWriter.Header().
		Set("Content-Type", "application/json")

	presenter.responseWriter.WriteHeader(http.StatusOK)

	response := struct {
		Result *models.Event `json:"response"`
	}{
		Result: event,
	}

	encodedResponse, encodedError := json.Marshal(response)
	if encodedError != nil {
		log.Println(encodedError)
	}

	if _, err := presenter.responseWriter.Write(encodedResponse); err != nil {
		log.Println(err)
	}
}
