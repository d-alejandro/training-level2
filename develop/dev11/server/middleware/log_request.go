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

	requestHeader, encodedHeaderError := receiver.encodeRequest(request.Header)
	if encodedHeaderError != nil {
		log.Println("marshal error:", encodedHeaderError)
		return
	}

	if request.Form == nil {
		if err := request.ParseForm(); err != nil {
			log.Println("parse error:", err)
			return
		}
	}

	requestForm, encodedFormError := receiver.encodeRequest(request.Form)
	if encodedFormError != nil {
		log.Println("marshal error:", encodedFormError)
		return
	}

	log.Printf(
		"\nRequest: %s %s\nRemoteAddress: %s\nHeaders:\n%s\nParameters:\n%s\nTime: %s",
		request.Method,
		request.RequestURI,
		request.RemoteAddr,
		requestHeader,
		requestForm,
		duration,
	)
}

func (receiver *LogRequest) encodeRequest(request any) (string, error) {
	data, err := json.MarshalIndent(request, "", "  ")

	if err != nil {
		return "", err
	}

	return string(data), nil
}
