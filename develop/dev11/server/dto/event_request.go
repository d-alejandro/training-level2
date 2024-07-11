package dto

/*
EventRequestDTO structure
*/
type EventRequestDTO struct {
	name string
	date string
}

/*
NewEventRequestDTO constructor
*/
func NewEventRequestDTO(name, date string) *EventRequestDTO {
	return &EventRequestDTO{
		name: name,
		date: date,
	}
}

/*
GetName method
*/
func (receiver *EventRequestDTO) GetName() string {
	return receiver.name
}

/*
GetDate method
*/
func (receiver *EventRequestDTO) GetDate() string {
	return receiver.date
}
