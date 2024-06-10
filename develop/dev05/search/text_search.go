package search

import (
	"slices"
	"strings"
)

type TextSearch struct {
	dto *TextSearchFlagDTO
}

func NewTextSearch(dto *TextSearchFlagDTO) *TextSearch {
	return &TextSearch{dto}
}

func (receiver *TextSearch) Search(pattern string, rows []string) []string {
	var outputSlice []string

	foundRowMap := receiver.findRows(pattern, rows)
	outputMethodService := NewOutputMethodService(rows, foundRowMap)

	switch {
	case receiver.dto.RowsAfterFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsAfterFlag(receiver.dto.RowsAfterFlag)
	case receiver.dto.RowsBeforeFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsBeforeFlag(receiver.dto.RowsBeforeFlag)
	case receiver.dto.RowsContextFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsContextFlag(receiver.dto.RowsContextFlag)
	}

	return receiver.compactAndReplaceSlice(outputSlice)
}

func (receiver *TextSearch) findRows(pattern string, rows []string) map[int]string {
	searchRows := make(map[int]string)

	for key, row := range rows {
		if strings.Contains(row, pattern) {
			colorText := "\u001B[31m" + pattern + "\u001B[0m"
			searchRows[key] = strings.ReplaceAll(row, pattern, colorText)
		}
	}

	return searchRows
}

func (receiver *TextSearch) compactAndReplaceSlice(rows []string) []string {
	tempSlice := slices.CompactFunc(rows, func(a string, b string) bool {
		if a == "" && a == b {
			return true
		}
		return false
	})

	firstIndex := 0

	if tempSlice[firstIndex] == "" {
		firstIndex++
	}

	lastIndex := len(tempSlice)

	if tempSlice[lastIndex-1] == "" {
		lastIndex--
	}

	tempSlice = slices.Clone(tempSlice[firstIndex:lastIndex])

	for key, value := range tempSlice {
		if value == "" {
			tempSlice[key] = "\u001B[34m--\u001B[0m"
		}
	}

	return tempSlice
}
