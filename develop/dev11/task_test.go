package main

import (
	"d-alejandro/training-level2/develop/dev11/server"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestEndpoints(t *testing.T) {
	tests := []struct {
		name                      string
		method                    string
		url                       string
		requestBody               func() *strings.Reader
		expectedStatus            int
		expectedBodyRootKey       string
		expectedHeaderContentType string
	}{
		{
			name:   "POST /create_event",
			method: http.MethodPost,
			url:    "/create_event",
			requestBody: func() *strings.Reader {
				data := url.Values{}
				data.Set("name", "Test Name")
				data.Set("date", "2024-07-11")
				return strings.NewReader(data.Encode())
			},
			expectedBodyRootKey:       `{"result":`,
			expectedStatus:            http.StatusOK,
			expectedHeaderContentType: "application/json",
		},
		{
			name:   "Bad Request POST /create_event",
			method: http.MethodPost,
			url:    "/create_event",
			requestBody: func() *strings.Reader {
				data := url.Values{}
				data.Set("name", "Test Name")
				data.Set("date", "2024-07-33")
				return strings.NewReader(data.Encode())
			},
			expectedBodyRootKey:       `{"error":`,
			expectedStatus:            http.StatusBadRequest,
			expectedHeaderContentType: "application/json",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpServer := server.NewHTTPServer()
			handler := httpServer.InitRequestHandler()

			responseRecorder := httptest.NewRecorder()
			request, _ := http.NewRequest(test.method, test.url, test.requestBody())
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			handler.ServeHTTP(responseRecorder, request)

			result := responseRecorder.Result()

			body, err := io.ReadAll(result.Body)
			if err != nil {
				t.Error(err)
				return
			}

			if result.StatusCode != test.expectedStatus {
				t.Errorf("expected status %d, actual %d", test.expectedStatus, responseRecorder.Code)
				return
			}

			resultContentType := result.Header.Get("Content-Type")

			if resultContentType != test.expectedHeaderContentType {
				t.Errorf(
					"expected header Content-Type %s, actual header Content-Type %s",
					test.expectedHeaderContentType,
					resultContentType,
				)
				return
			}

			bodyResult := string(body)

			if !strings.HasPrefix(bodyResult, test.expectedBodyRootKey) {
				t.Errorf(
					"expect response body to contain %s, actual response body to contain %s",
					test.expectedBodyRootKey,
					string(body),
				)
			}
		})
	}
}
