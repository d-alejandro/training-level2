package helpers

import (
	"errors"
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

type NTPAdapter struct {
}

func NewNTPAdapter() *NTPAdapter {
	return &NTPAdapter{}
}

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
