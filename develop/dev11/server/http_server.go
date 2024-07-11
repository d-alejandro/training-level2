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
	handler := receiver.InitRequestHandler()

	err := http.ListenAndServe(receiver.getNetworkAddress(), handler)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (receiver *HTTPServer) InitRequestHandler() http.Handler {
	serveMux := http.NewServeMux()
	BindRouteHandlers(serveMux, bindings.NewHandlerBinding())
	return receiver.bindMiddleware(serveMux)
}

func (receiver *HTTPServer) bindMiddleware(serveMux *http.ServeMux) http.Handler {
	handler := middleware.NewLogRequest(serveMux)
	handler = middleware.NewPanicRecovery(handler)
	return handler
}

func (receiver *HTTPServer) getNetworkAddress() string {
	httpConfigs := GetConfigs()["http"].(map[string]string)
	return net.JoinHostPort(httpConfigs["host"], httpConfigs["port"])
}
