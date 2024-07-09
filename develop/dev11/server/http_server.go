package server

import (
	"d-alejandro/training-level2/develop/dev11/server/bindings"
	"d-alejandro/training-level2/develop/dev11/server/middleware"
	"fmt"
	"net"
	"net/http"
	"os"
)

type HTTPServer struct {
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (receiver *HTTPServer) ListenAndServe() {
	serveMux := http.NewServeMux()

	handlerBinding := bindings.NewHandlerBinding()
	InitRoutes(serveMux, handlerBinding)

	handler := middleware.NewLogRequest(serveMux)
	handler = middleware.NewPanicRecovery(handler)

	httpConfigs := GetConfigs()["http"].(map[string]string)
	address := net.JoinHostPort(httpConfigs["host"], httpConfigs["port"])

	if err := http.ListenAndServe(address, handler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
