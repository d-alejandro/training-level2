package main

import (
	"d-alejandro/training-level2/develop/dev08/cmd"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestHandlerExecute(t *testing.T) {
	currentDirectory, directoryError := os.Getwd()

	if directoryError != nil {
		fmt.Println("Error:", directoryError.Error())
		return
	}

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
		{
			name:             "test cd",
			inputCommandRow:  `cd cmd`,
			expectedResponse: currentDirectory + "/cmd",
			expectedError:    nil,
		},
		{
			name:             "test pwd",
			inputCommandRow:  `pwd`,
			expectedResponse: currentDirectory + "\n",
			expectedError:    nil,
		},
		{
			name:             "test echo",
			inputCommandRow:  `echo 'qwerty ty jjo'`,
			expectedResponse: "qwerty ty jjo\n",
			expectedError:    nil,
		},
		{
			name:             "test kill",
			inputCommandRow:  `kill -l SIGQUIT`,
			expectedResponse: "3\n",
			expectedError:    nil,
		},
		{
			name:             "test ps",
			inputCommandRow:  `ps -C ps -o comm`,
			expectedResponse: "COMMAND\nps\n",
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

			if test.name == "test fork/exec command" {
				if pid, err := strconv.Atoi(response); err == nil && pid > 1 {
					return
				}
				t.Errorf("handler.Execute() = '%v', expected '%v'", response, test.expectedResponse)
				return
			} else if test.name == "test cd" {
				dir, dirError := os.Getwd()

				if dirError != nil {
					t.Error(dirError.Error())
					return
				}

				if response != "Ok" || dir != test.expectedResponse {
					t.Errorf(
						"handler.Execute() = '%v', current directory '%v', expected '%v'",
						response,
						dir,
						test.expectedResponse,
					)
					return
				}
				return
			}

			if response != test.expectedResponse {
				t.Errorf("handler.Execute() = '%v', expected '%v'", response, test.expectedResponse)
			}
		})
	}

	close(forkExecResultChannel)
}
