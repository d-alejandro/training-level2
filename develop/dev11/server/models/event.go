package models

import "time"

/*
Event structure
*/
type Event struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
