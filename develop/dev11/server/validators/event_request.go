package validators

import (
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

func (receiver *EventRequestValidator) Validate(request *http.Request) error {
	if err := request.ParseForm(); err != nil {
		return err
	}

	receiver.Name = request.FormValue("name")
	receiver.Date = request.FormValue("date")

	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(receiver)
}
