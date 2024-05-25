package test

import (
	"d-alejandro/training-level2/develop/dev01/internal/helpers"
	"testing"
	"time"
)

func TestSuccessfulGetCurrentTimeOfNTPAdapter(t *testing.T) {
	ntpAdapter := helpers.NewNTPAdapter()

	const NTPServerAddress = "0.beevik-ntp.pool.ntp.org"
	ntpTime, err := ntpAdapter.GetCurrentTime(NTPServerAddress)

	if err != nil {
		actualErrorMessage := err.Error()
		t.Errorf(`got "%s" but expected "nil"`, actualErrorMessage)
	}

	var emptyTime time.Time

	const TimeLayout = "2006-01-02 15:04:05.000000000"

	emptyTimeFormated := emptyTime.Format(TimeLayout)
	ntpTimeFormated := ntpTime.Format(TimeLayout)

	if emptyTimeFormated == ntpTimeFormated {
		t.Errorf(`got "%s" but expected not empty time`, ntpTimeFormated)
	}
}

func TestFailedGetCurrentTimeOfNTPAdapter(t *testing.T) {
	tests := []struct {
		name                 string
		ntpServerArg         string
		expectedErrorMessage string
	}{
		{
			name:                 "ntpServer argument is empty",
			ntpServerArg:         "",
			expectedErrorMessage: "address string is empty",
		},
		{
			name:                 "ntpServer argument is 0",
			ntpServerArg:         "0",
			expectedErrorMessage: "lookup 0: no such host",
		},
	}

	ntpAdapter := helpers.NewNTPAdapter()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ntpAdapter.GetCurrentTime(test.ntpServerArg)

			if err != nil {
				actualErrorMessage := err.Error()

				if test.expectedErrorMessage != actualErrorMessage {
					t.Errorf(`got "%s" but expected "%s"`, actualErrorMessage, test.expectedErrorMessage)
				}
			}
		})
	}
}
