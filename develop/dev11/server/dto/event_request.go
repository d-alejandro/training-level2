package dto

type EventRequestDTO struct {
	name string
	date string
}

func NewEventRequestDTO(name, date string) *EventRequestDTO {
	return &EventRequestDTO{
		name: name,
		date: date,
	}
}

func (receiver *EventRequestDTO) GetName() string {
	return receiver.name
}

func (receiver *EventRequestDTO) GetDate() string {
	return receiver.date
}
