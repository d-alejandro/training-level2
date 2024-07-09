package validators

import (
	"d-alejandro/training-level2/develop/dev11/server/dto"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type EventRequestValidator struct {
	Name string `json:"name" validate:"required"`
	Date string `json:"date" validate:"required,datetime=2006-01-02"`
}

func NewEventRequestValidator() *EventRequestValidator {
	return &EventRequestValidator{}
}

func (receiver *EventRequestValidator) Validate(request *http.Request) (*dto.EventRequestDTO, error) {
	if err := request.ParseForm(); err != nil {
		return nil, err
	}

	receiver.Name = request.FormValue("name")
	receiver.Date = request.FormValue("date")

	eventRequestDTO := dto.NewEventRequestDTO(receiver.Name, receiver.Date)

	validate := validator.New(validator.WithRequiredStructEnabled())

	return eventRequestDTO, validate.Struct(receiver)
}
