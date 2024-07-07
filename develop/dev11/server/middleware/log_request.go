package middleware

import (
	"encoding/json"
	"fmt"
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
	start := time.Now()

	receiver.nextHandler.ServeHTTP(responseWriter, request)

	requestHeader, marshalHeaderError := json.MarshalIndent(request.Header, "", "  ")
	if marshalHeaderError != nil {
		fmt.Println("marshal error:", marshalHeaderError)
		return
	}

	parseError := request.ParseForm()
	if parseError != nil {
		fmt.Println("parse error:", parseError)
		return
	}

	requestForm, marshalFormError := json.MarshalIndent(request.Form, "", "  ")
	if marshalFormError != nil {
		fmt.Println("marshal error:", marshalFormError)
		return
	}

	log.Printf(
		"\nRequest: %s %s\nRemoteAddress: %s\nHeaders:\n%s\nParameters:\n%s\nTimeoutDuration: %s",
		request.Method,
		request.RequestURI,
		request.RemoteAddr,
		string(requestHeader),
		string(requestForm),
		time.Since(start),
	)
}
