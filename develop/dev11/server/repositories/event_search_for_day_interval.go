package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"slices"
	"time"
)

/*
EventSearchForDayIntervalRepository structure
*/
type EventSearchForDayIntervalRepository struct {
	dbConnection database.CacheContract
}

/*
NewEventSearchForDayIntervalRepository constructor
*/
func NewEventSearchForDayIntervalRepository(dbConnection database.CacheContract) *EventSearchForDayIntervalRepository {
	return &EventSearchForDayIntervalRepository{dbConnection}
}

/*
Make method
*/
func (receiver *EventSearchForDayIntervalRepository) Make(startDate, endDate *time.Time) ([]*models.Event, error) {
	events, loadError := receiver.dbConnection.LoadEvents()
	if loadError != nil {
		return nil, loadError
	}

	var result []*models.Event

	for _, event := range events {
		if (*startDate == event.Date) || (startDate.Before(event.Date) && endDate.After(event.Date)) {
			result = append(result, event)
		}
	}

	slices.SortStableFunc(result, func(a, b *models.Event) int {
		return a.Date.Compare(b.Date)
	})

	return result, nil
}
