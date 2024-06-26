package main

import (
	"d-alejandro/training-level2/develop/dev09/getter"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWebGetter_Execute(t *testing.T) {
	currentDirectory, directoryError := os.Getwd()

	if directoryError != nil {
		fmt.Println("Error:", directoryError.Error())
		return
	}

	tests := []struct {
		name                 string
		levelMaxFlag         int
		url                  string
		expectedFilePath     string
		expectedErrorMessage string
	}{
		{
			name:                 "test with -l 1 https://www.shakira.com/",
			levelMaxFlag:         1,
			url:                  "https://www.shakira.com/",
			expectedFilePath:     "/www.shakira.com/index.html",
			expectedErrorMessage: "",
		},
		{
			name:                 "test with -l 3 http://localhost/",
			levelMaxFlag:         3,
			url:                  "http://localhost/",
			expectedFilePath:     "",
			expectedErrorMessage: `Get "http://localhost/": dial tcp 127.0.0.1:80: connect: connection refused`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			webGetter := getter.NewWebGetter(test.levelMaxFlag)
			actualError := webGetter.Execute(test.url)

			if actualError != nil && !strings.Contains(actualError.Error(), test.expectedErrorMessage) {
				t.Errorf(
					"webGetter.Execute(%v) error = '%v', expected error '%v'",
					test.url,
					actualError,
					test.expectedErrorMessage,
				)
				return
			}

			if test.expectedFilePath == "" {
				return
			}

			filePath := currentDirectory + test.expectedFilePath

			if _, err := os.Stat(filePath); err != nil {
				t.Errorf("webGetter.Execute(%v), expected file: '%v'", test.url, filePath)
			}

			filePath = filepath.Clean(filePath + "/../")

			if err := os.RemoveAll(filePath); err != nil {
				fmt.Println("Error:", err.Error())
			}
		})
	}
}
