package main

import (
	"d-alejandro/training-level2/develop/dev08/cmd"
	"strconv"
	"testing"
)

func TestHandlerExecute(t *testing.T) {
	tests := []struct {
		name             string
		inputCommandRow  string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "test fork/exec command",
			inputCommandRow:  `echo test &`,
			expectedResponse: "integer PID",
			expectedError:    nil,
		},
		{
			name:             "test pipes: cmd1 | cmd2 | ... | cmdN",
			inputCommandRow:  `echo "RR www qq ff ll" | tr -d 'wl' | wc -w | tr 3 7`,
			expectedResponse: "7\n",
			expectedError:    nil,
		},
	}

	forkExecResultChannel := make(chan string, 1)
	handler := cmd.NewHandler(forkExecResultChannel)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, responseError := handler.Execute(test.inputCommandRow)

			if responseError != nil && responseError.Error() != test.expectedError.Error() {
				t.Errorf("handler.Execute() error = '%v', expected error '%v'", responseError, test.expectedError)
				return
			}

			if test.expectedResponse == "integer PID" {
				if pid, err := strconv.Atoi(response); err == nil && pid > 1 {
					return
				}
				t.Errorf("handler.Execute() = '%v', expected '%v'", response, test.expectedResponse)
				return
			}

			if response != test.expectedResponse {
				t.Errorf("handler.Execute() = '%v', expected '%v'", response, test.expectedResponse)
			}
		})
	}

	close(forkExecResultChannel)
}
