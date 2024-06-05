package sorting

import (
	"regexp"
	"slices"
	"strings"
)

type FileSorting struct {
	orderFlagDTO *FlagDTO
}

func NewFileSorting(orderFlagDTO *FlagDTO) *FileSorting {
	return &FileSorting{orderFlagDTO}
}

func (receiver *FileSorting) Sort(array []string) []string {
	sortedArray := make([][]string, len(array))

	regExpr := regexp.MustCompile(`\S+\s*`)

	for key, value := range array {
		sortedArray[key] = receiver.splitAndTrimRow(regExpr, value)
	}

	slices.SortFunc(sortedArray, func(a, b []string) int {
		firstValue := strings.ToLower(a[receiver.orderFlagDTO.ColumnNumberFlag])
		secondValue := strings.ToLower(b[receiver.orderFlagDTO.ColumnNumberFlag])

		if result := strings.Compare(firstValue, secondValue); result != 0 {
			return result
		}

		return strings.Compare(b[receiver.orderFlagDTO.ColumnNumberFlag], a[receiver.orderFlagDTO.ColumnNumberFlag])
	})

	joinedArray := make([]string, len(sortedArray))

	for index, value := range sortedArray {
		joinedArray[index] = strings.Join(value, " ")
	}

	return joinedArray
}

func (receiver *FileSorting) splitAndTrimRow(regexp *regexp.Regexp, row string) []string {
	array := regexp.FindAllString(row, -1)

	var functionTrim func(key int, value string)

	if receiver.orderFlagDTO.EndingSpaceFlag {
		functionTrim = func(key int, value string) {
			array[key] = strings.TrimRight(value, " ")
		}
	} else {
		functionTrim = func(key int, value string) {
			array[key] = strings.TrimSuffix(value, " ")
		}
	}

	for key, value := range array {
		functionTrim(key, value)
	}

	return array
}
