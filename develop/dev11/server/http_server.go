package server

import (
	"d-alejandro/training-level2/develop/dev11/server/handlers"
	"net"
	"net/http"
)

type HTTPServer struct {
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (receiver *HTTPServer) ListenAndServe() {
	serveMux := http.NewServeMux()

	serveMux.Handle("POST /api/events/create", handlers.NewEventCreationHandler())
	serveMux.Handle("POST /api/events/update", handlers.NewEventUpdateHandler())

	httpConfigs := GetConfigs()["http"].(map[string]string)

	address := net.JoinHostPort(httpConfigs["host"], httpConfigs["port"])

	http.ListenAndServe(address, serveMux)
}
