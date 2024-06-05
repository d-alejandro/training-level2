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

	slices.SortFunc(sortedArray, receiver.runSortFunc)

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

func (receiver *FileSorting) runSortFunc(a, b []string) int {
	var columnStart int

	if receiver.orderFlagDTO.ColumnNumberFlag > 0 {
		columnStart = receiver.orderFlagDTO.ColumnNumberFlag - 1
	}

	firstValue := strings.ToLower(a[columnStart])
	secondValue := strings.ToLower(b[columnStart])

	sortMethod := NewSortMethod(receiver.orderFlagDTO)

	if result := sortMethod.Execute(firstValue, secondValue); result != 0 {
		if receiver.orderFlagDTO.DescendingOrderFlag {
			return -1 * result
		}
		return result
	}

	if receiver.orderFlagDTO.DescendingOrderFlag {
		return strings.Compare(a[columnStart], b[columnStart])
	}

	return strings.Compare(b[columnStart], a[columnStart])
}
