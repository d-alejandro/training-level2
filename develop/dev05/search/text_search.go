package search

import "strings"

type TextSearch struct {
	dto        *TextSearchFlagDTO
	colorToRed func(text string) string
}

func NewTextSearch(dto *TextSearchFlagDTO) *TextSearch {
	funcColorToRed := func(text string) string {
		return "\u001B[31m" + text + "\u001B[0m"
	}

	return &TextSearch{
		dto:        dto,
		colorToRed: funcColorToRed,
	}
}

func (receiver *TextSearch) Search(pattern string, rows []string) []string {
	var searchRows []string

	for _, row := range rows {
		if strings.Contains(row, pattern) {
			formattedRow := strings.ReplaceAll(row, pattern, receiver.colorToRed(pattern))
			searchRows = append(searchRows, formattedRow)
		}
	}

	return searchRows
}
