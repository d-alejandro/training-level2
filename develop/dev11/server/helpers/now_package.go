package helpers

import (
	"github.com/jinzhu/now"
	"time"
)

/*
GetNow function
*/
func GetNow(date *time.Time) (*now.Now, error) {
	const LocationMoscow = "Europe/Moscow"
	location, err := time.LoadLocation(LocationMoscow)
	if err != nil {
		return nil, err
	}

	config := &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02 15:04:05"},
	}

	return config.With(*date), nil
}
