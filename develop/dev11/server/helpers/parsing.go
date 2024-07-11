package helpers

import "time"

/*
ParseDate function
*/
func ParseDate(date string) (*time.Time, error) {
	const layout = "2006-01-02"

	parsedDate, err := time.Parse(layout, date)

	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}
