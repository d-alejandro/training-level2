package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorPresenter struct {
	responseWriter http.ResponseWriter
}

func NewErrorPresenter(responseWriter http.ResponseWriter) *ErrorPresenter {
	return &ErrorPresenter{responseWriter}
}

func (presenter *ErrorPresenter) Present(statusCode int, errorMessage any) {
	presenter.responseWriter.Header().
		Set("Content-Type", "application/json")

	presenter.responseWriter.WriteHeader(statusCode)

	message := fmt.Sprint(errorMessage)
	status := http.StatusText(statusCode)
	errorResponse := NewErrorResponse(message, status)

	response, encodedError := json.Marshal(errorResponse)
	if encodedError != nil {
		panic(encodedError)
	}

	if _, err := presenter.responseWriter.Write(response); err != nil {
		panic(err)
	}
}
