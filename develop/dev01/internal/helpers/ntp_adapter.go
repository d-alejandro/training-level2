package helpers

import (
	"errors"
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

/*
NTPAdapter is an adapter for the NTP library
*/
type NTPAdapter struct {
}

/*
NewNTPAdapter is the NTPAdapter constructor
*/
func NewNTPAdapter() *NTPAdapter {
	return &NTPAdapter{}
}

/*
GetCurrentTime is an action for NTPAdapter
*/
func (receiver *NTPAdapter) GetCurrentTime(ntpServer string) (time time.Time, err error) {
	defer func() {
		if recoveryError := recover(); recoveryError != nil {
			err = receiver.convertToError(recoveryError)
		}
	}()

	return ntp.Time(ntpServer)
}

func (receiver *NTPAdapter) convertToError(recoveryError any) error {
	switch errorType := recoveryError.(type) {
	case error:
		return errorType
	case string:
		return errors.New(errorType)
	default:
		return errors.New(fmt.Sprint(errorType))
	}
}
