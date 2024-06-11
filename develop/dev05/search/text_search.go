package search

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

/*
TextSearch structure
*/
type TextSearch struct {
	dto *TextSearchFlagDTO
}

/*
NewTextSearch constructor
*/
func NewTextSearch(dto *TextSearchFlagDTO) *TextSearch {
	return &TextSearch{dto}
}

/*
Search method
*/
func (receiver *TextSearch) Search(pattern string, rows []string) []string {
	var outputSlice []string

	foundRowMap := receiver.findRows(pattern, rows)
	outputMethodService := NewOutputMethodService(rows, foundRowMap)

	switch {
	case receiver.dto.RowsAfterFlag > 0 && receiver.dto.RowsBeforeFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsAfterAndBeforeFlag(
			receiver.dto.RowsAfterFlag,
			receiver.dto.RowsBeforeFlag,
		)
		receiver.addLineNumIfFlagSet(outputSlice, foundRowMap)
		outputSlice = receiver.compactAndReplaceSlice(outputSlice)
	case receiver.dto.RowsAfterFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsAfterFlag(receiver.dto.RowsAfterFlag)
		receiver.addLineNumIfFlagSet(outputSlice, foundRowMap)
		outputSlice = receiver.compactAndReplaceSlice(outputSlice)
	case receiver.dto.RowsBeforeFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsBeforeFlag(receiver.dto.RowsBeforeFlag)
		receiver.addLineNumIfFlagSet(outputSlice, foundRowMap)
		outputSlice = receiver.compactAndReplaceSlice(outputSlice)
	case receiver.dto.RowsContextFlag > 0:
		outputSlice = outputMethodService.ExecuteForRowsContextFlag(receiver.dto.RowsContextFlag)
		receiver.addLineNumIfFlagSet(outputSlice, foundRowMap)
		outputSlice = receiver.compactAndReplaceSlice(outputSlice)
	case receiver.dto.CountFlag:
		resultRowCount := strconv.Itoa(len(foundRowMap))
		return []string{resultRowCount}
	case receiver.dto.InvertFlag:
		outputSlice = outputMethodService.ExecuteForInvertFlag()
		receiver.addLineNumIfFlagSet(outputSlice, foundRowMap)
		outputSlice = receiver.removeEmptyLines(outputSlice)
	default:
		outputSlice = outputMethodService.ExecuteWithoutFlags()
		receiver.addLineNumIfFlagSet(outputSlice, foundRowMap)
		outputSlice = receiver.removeEmptyLines(outputSlice)
	}

	return outputSlice
}

func (receiver *TextSearch) findRows(pattern string, rows []string) map[int]string {
	searchRows := make(map[int]string)

	if receiver.dto.FixedFlag {
		pattern = regexp.QuoteMeta(pattern)
	}

	if receiver.dto.IgnoreCaseFlag {
		pattern = "(?i)" + pattern
	}

	regExpr := regexp.MustCompile(pattern)

	for key, row := range rows {
		result := regExpr.FindAllString(row, -1)

		if result == nil {
			continue
		}

		for _, value := range result {
			colorText := "\u001B[31m" + value + "\u001B[0m"
			searchRows[key] = strings.ReplaceAll(row, value, colorText)
		}
	}

	return searchRows
}

func (receiver *TextSearch) addLineNumIfFlagSet(outputSlice []string, foundRowMap map[int]string) {
	if !receiver.dto.LineNumFlag {
		return
	}

	for key, row := range outputSlice {
		if row == "" {
			continue
		}

		var symbol string

		if _, ok := foundRowMap[key]; ok {
			symbol = ":"
		} else {
			symbol = "-"
		}

		outputSlice[key] = fmt.Sprintf("\u001B[32m%d\u001B[0m\u001B[34m%s\u001B[0m%s", key+1, symbol, row)
	}
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

func (receiver *TextSearch) removeEmptyLines(outputSlice []string) []string {
	return slices.DeleteFunc(outputSlice, func(row string) bool {
		return row == ""
	})
}
