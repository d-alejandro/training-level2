package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"slices"
	"strings"
	"time"
)

type EventShowForDayRepository struct {
	dbConnection database.CacheContract
}

func NewEventShowForDayRepository(dbConnection database.CacheContract) *EventShowForDayRepository {
	return &EventShowForDayRepository{dbConnection}
}

func (receiver *EventShowForDayRepository) Make(date *time.Time) ([]*models.Event, error) {
	events, loadError := receiver.dbConnection.LoadEvents()
	if loadError != nil {
		return nil, loadError
	}

	var result []*models.Event

	for _, event := range events {
		if *date == event.Date {
			result = append(result, event)
		}
	}

	slices.SortStableFunc(result, func(a, b *models.Event) int {
		return strings.Compare(a.Name, b.Name)
	})

	return result, nil
}
