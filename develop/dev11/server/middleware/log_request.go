package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type LogRequest struct {
	nextHandler http.Handler
}

func NewLogRequest(handler http.Handler) *LogRequest {
	return &LogRequest{
		nextHandler: handler,
	}
}

func (receiver *LogRequest) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	timeStart := time.Now()

	receiver.nextHandler.ServeHTTP(responseWriter, request)

	duration := time.Since(timeStart)

	requestHeader, marshalHeaderError := json.MarshalIndent(request.Header, "", "  ")
	if marshalHeaderError != nil {
		log.Println("marshal error:", marshalHeaderError)
		return
	}

	parseError := request.ParseForm()
	if parseError != nil {
		log.Println("parse error:", parseError)
		return
	}

	requestForm, marshalFormError := json.MarshalIndent(request.Form, "", "  ")
	if marshalFormError != nil {
		log.Println("marshal error:", marshalFormError)
		return
	}

	log.Printf(
		"\nRequest: %s %s\nRemoteAddress: %s\nHeaders:\n%s\nParameters:\n%s\nTime: %s",
		request.Method,
		request.RequestURI,
		request.RemoteAddr,
		string(requestHeader),
		string(requestForm),
		duration,
	)
}
