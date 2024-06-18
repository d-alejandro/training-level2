package main

import (
	"testing"
	"time"
)

func Test_mergeChannels(t *testing.T) {
	signalTest := func(after time.Duration) <-chan any {
		channel := make(chan interface{})
		go func() {
			defer close(channel)
			time.Sleep(after)
		}()
		return channel
	}

	tests := []struct {
		name                            string
		inputChannels                   []<-chan any
		expectedDurationLessThanOrEqual time.Duration
		expectedResponse                any
	}{
		{
			name: "done after 1 s",
			inputChannels: []<-chan any{
				signalTest(3 * time.Second),
				signalTest(2 * time.Second),
				signalTest(1 * time.Second),
			},
			expectedDurationLessThanOrEqual: 2 * time.Second,
			expectedResponse:                nil,
		},
		{
			name: "done after 3 s",
			inputChannels: []<-chan any{
				signalTest(5 * time.Second),
				signalTest(3 * time.Second),
				signalTest(7 * time.Second),
			},
			expectedDurationLessThanOrEqual: 4 * time.Second,
			expectedResponse:                nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := time.Now()

			select {
			case response := <-mergeChannels(test.inputChannels...):
				if response != test.expectedResponse {
					t.Errorf("Expected %v, got %v", test.expectedResponse, response)
				}
			case <-time.After(test.expectedDurationLessThanOrEqual):
				t.Errorf(
					"Test timed out.\nExpected duration: %s.\nActual duration: %s.",
					test.expectedDurationLessThanOrEqual,
					time.Since(start),
				)
			}
		})
	}
}
