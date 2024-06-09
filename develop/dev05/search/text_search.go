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
	rowLength := len(rows)
	outputSlice := make([]string, rowLength)

	foundRowMap := receiver.findRows(pattern, rows)

	switch {
	case receiver.dto.RowCountAfter > 0:
		for key := range outputSlice {
			if value, ok := foundRowMap[key]; ok {
				outputSlice[key] = value

				lastIndex := key + 1 + receiver.dto.RowCountAfter

				if lastIndex > rowLength {
					lastIndex = rowLength
				}

				for index := key + 1; index < lastIndex; index++ {
					outputSlice[index] = rows[index]
				}
			}
		}
	}

	return outputSlice
}

func (receiver *TextSearch) findRows(pattern string, rows []string) map[int]string {
	searchRows := make(map[int]string)

	for key, row := range rows {
		if strings.Contains(row, pattern) {
			searchRows[key] = strings.ReplaceAll(row, pattern, receiver.colorToRed(pattern))
		}
	}

	return searchRows
}
