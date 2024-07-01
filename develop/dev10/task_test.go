package main

import (
	"d-alejandro/training-level2/develop/dev10/socket"
	"errors"
	"testing"
	"time"
)

func TestClient_Send(t *testing.T) {
	const (
		TestNetworkProtocolTCP = "tcp"
		TestNetworkHost        = "127.0.0.1"
		TestNetworkPort        = "8083"
		TestNetworkAddress     = TestNetworkHost + ":" + TestNetworkPort
		TestTimeoutDuration    = 1 * time.Second
	)

	server, serverError := socket.NewServer(TestNetworkProtocolTCP, TestNetworkAddress)
	if serverError != nil {
		t.Fatal(serverError)
	}

	client, clientError := socket.NewClient(TestNetworkHost, TestNetworkPort, TestTimeoutDuration)
	if clientError != nil {
		t.Fatal(clientError)
	}

	tests := []struct {
		name             string
		inputMessage     string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "Привет, мир!",
			inputMessage:     "Привет, мир!",
			expectedResponse: "read length: 21, size: 16",
			expectedError:    nil,
		},
		{
			name:             "empty message",
			inputMessage:     "",
			expectedResponse: "",
			expectedError:    errors.New("message is empty"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualResponse, actualError := client.Send(test.inputMessage)

			if !isEqualErrors(actualError, test.expectedError) {
				t.Errorf("client.Send() error = '%v', expected error '%v'", actualError, test.expectedError)
				return
			}

			if actualResponse != test.expectedResponse {
				t.Errorf("client.Send() = '%v', expected '%v'", actualResponse, test.expectedResponse)
			}
		})
	}

	client.Stop()

	if err := server.Stop(); err != nil {
		t.Fatal(err)
	}
}

func isEqualErrors(actual error, expected error) bool {
	if actual == nil && expected == nil {
		return true
	} else if actual != nil && expected != nil {
		if actual.Error() == expected.Error() {
			return true
		}
	}
	return false
}
