package server

import (
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

	InitRoutes(serveMux)

	logRequest := middleware.NewLogRequest(serveMux)

	httpConfigs := GetConfigs()["http"].(map[string]string)
	address := net.JoinHostPort(httpConfigs["host"], httpConfigs["port"])

	err := http.ListenAndServe(address, logRequest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
