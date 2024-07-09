package middleware

import (
	"d-alejandro/training-level2/develop/dev11/server/presenters"
	"log"
	"net/http"
	"runtime/debug"
)

type PanicRecovery struct {
	nextHandler http.Handler
}

func NewPanicRecovery(handler http.Handler) http.Handler {
	return &PanicRecovery{
		nextHandler: handler,
	}
}

func (receiver *PanicRecovery) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			errorPresenter := presenters.NewErrorPresenter(responseWriter)
			errorPresenter.Present(http.StatusInternalServerError, err)

			log.Println(err)
			log.Println(string(debug.Stack()))
		}
	}()

	receiver.nextHandler.ServeHTTP(responseWriter, request)
}
