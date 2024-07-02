package main

import (
	"d-alejandro/training-level2/develop/dev10/socket"
	"testing"
	"time"
)

func TestTimeoutClientSend(t *testing.T) {
	tests := []struct {
		name                     string
		inputHost                string
		inputPort                string
		inputTimeoutDurationFlag time.Duration
		expectedDuration         time.Duration
	}{
		{
			name:                     "timeout 3 seconds",
			inputHost:                "mail.ru",
			inputPort:                "8080",
			inputTimeoutDurationFlag: 3 * time.Second,
			expectedDuration:         4 * time.Second,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			done := make(chan struct{})

			start := time.Now()

			go func() {
				_, _ = socket.NewClient(test.inputHost, test.inputPort, test.inputTimeoutDurationFlag)
				close(done)
			}()

			select {
			case <-time.After(test.expectedDuration):
				select {
				case <-time.After(test.expectedDuration):
				case <-done:
				}

				t.Errorf(
					"Test timeout duration.\nInput: %s.\nExpected (less than or equal): %s.\nActual: %s.",
					test.inputTimeoutDurationFlag,
					test.expectedDuration,
					time.Since(start),
				)
			case <-done:
			}
		})
	}
}
