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

	jsonModel, encodedError := json.Marshal(event)
	if encodedError != nil {
		log.Println(encodedError)
	}

	if _, err := presenter.responseWriter.Write(jsonModel); err != nil {
		log.Println(err)
	}
}
