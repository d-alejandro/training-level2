package presenters

import (
	"d-alejandro/training-level2/develop/dev11/server/presenters/resources"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
ErrorPresenter structure
*/
type ErrorPresenter struct {
	responseWriter http.ResponseWriter
}

/*
NewErrorPresenter constructor
*/
func NewErrorPresenter(responseWriter http.ResponseWriter) *ErrorPresenter {
	return &ErrorPresenter{responseWriter}
}

/*
Present method
*/
func (presenter *ErrorPresenter) Present(statusCode int, errorMessage any) {
	presenter.responseWriter.Header().
		Set("Content-Type", "application/json")

	presenter.responseWriter.WriteHeader(statusCode)

	message := fmt.Sprint(errorMessage)
	status := http.StatusText(statusCode)
	errorResponse := resources.NewErrorResponse(message, status)

	encodedErrorResponse, encodedError := json.Marshal(errorResponse)
	if encodedError != nil {
		log.Println(encodedError)
	}

	if _, err := presenter.responseWriter.Write(encodedErrorResponse); err != nil {
		log.Println(err)
	}
}
