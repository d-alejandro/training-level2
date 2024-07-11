package validators

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type EventDateRequestValidator struct {
	Date string `json:"date" validate:"required,datetime=2006-01-02"`
}

func NewEventDateRequestValidator() *EventDateRequestValidator {
	return &EventDateRequestValidator{}
}

func (receiver *EventDateRequestValidator) Validate(request *http.Request) (string, error) {
	if err := request.ParseForm(); err != nil {
		return "", err
	}

	receiver.Date = request.FormValue("date")

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(receiver); err != nil {
		return "", err
	}

	return receiver.Date, nil
}
