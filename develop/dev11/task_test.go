package main

import (
	"d-alejandro/training-level2/develop/dev11/server"
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestEndpoints(t *testing.T) {
	dbConnection := database.GetDatabaseConnection()

	id := "1"
	event := &models.Event{
		ID:   id,
		Name: "Test Event",
		Date: time.Now(),
	}
	_ = dbConnection.SetEvent(id, event)

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
		{
			name:   "POST /update_event/" + id,
			method: http.MethodPost,
			url:    "/update_event/" + id,
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
			name:   "POST /delete_event/" + id,
			method: http.MethodPost,
			url:    "/delete_event/" + id,
			requestBody: func() *strings.Reader {
				return strings.NewReader("")
			},
			expectedBodyRootKey:       `{"result":`,
			expectedStatus:            http.StatusOK,
			expectedHeaderContentType: "application/json",
		},
		{
			name:   "GET /events_for_day?date=2000-01-01",
			method: http.MethodGet,
			url:    "/events_for_day?date=2000-01-01",
			requestBody: func() *strings.Reader {
				return strings.NewReader("")
			},
			expectedBodyRootKey:       `{"result":`,
			expectedStatus:            http.StatusOK,
			expectedHeaderContentType: "application/json",
		},
		{
			name:   "GET /events_for_week?date=2000-01-01",
			method: http.MethodGet,
			url:    "/events_for_week?date=2000-01-01",
			requestBody: func() *strings.Reader {
				return strings.NewReader("")
			},
			expectedBodyRootKey:       `{"result":`,
			expectedStatus:            http.StatusOK,
			expectedHeaderContentType: "application/json",
		},
		{
			name:   "GET /events_for_month?date=2000-01-01",
			method: http.MethodGet,
			url:    "/events_for_month?date=2000-01-01",
			requestBody: func() *strings.Reader {
				return strings.NewReader("")
			},
			expectedBodyRootKey:       `{"result":`,
			expectedStatus:            http.StatusOK,
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

			resultBody := string(body)
			if !strings.HasPrefix(resultBody, test.expectedBodyRootKey) {
				t.Errorf(
					"expect response body to contain %s, actual response body to contain %s",
					test.expectedBodyRootKey,
					resultBody,
				)
			}
		})
	}
}
