package validators

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type EventRequestValidator struct {
	Name string `json:"name" validate:"required"`
	Date string `json:"date" validate:"required,datetime=2006-01-02"`
}

func NewEventRequestValidator(parsedRequest *http.Request) *EventRequestValidator {
	return &EventRequestValidator{
		Name: parsedRequest.FormValue("name"),
		Date: parsedRequest.FormValue("date"),
	}
}

func (receiver *EventRequestValidator) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(receiver)
}
