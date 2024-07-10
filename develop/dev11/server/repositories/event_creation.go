package repositories

import (
	"d-alejandro/training-level2/develop/dev11/server/database"
	"d-alejandro/training-level2/develop/dev11/server/models"
	"github.com/google/uuid"
	"time"
)

type EventCreationRepository struct {
	dbConnection database.CacheContract
}

func NewEventCreationRepository(dbConnection database.CacheContract) *EventCreationRepository {
	return &EventCreationRepository{dbConnection}
}

func (receiver *EventCreationRepository) Make(name string, date *time.Time) (*models.Event, error) {
	id := uuid.Must(uuid.NewRandom()).String()

	event := &models.Event{
		ID:   id,
		Name: name,
		Date: *date,
	}

	if err := receiver.dbConnection.SetEvent(id, event); err != nil {
		return nil, err
	}

	return event, nil
}
