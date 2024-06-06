package sorting

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

/*
ArraySorting structure - sort slice
*/
type ArraySorting struct {
	orderFlagDTO       *FlagDTO
	workingColumNumber int
}

/*
NewArraySorting is the ArraySorting constructor
*/
func NewArraySorting(orderFlagDTO *FlagDTO) *ArraySorting {
	workingColumNumber := 0

	if orderFlagDTO.ColumnNumberFlag > 0 {
		workingColumNumber = orderFlagDTO.ColumnNumberFlag - 1
	}

	return &ArraySorting{
		orderFlagDTO,
		workingColumNumber,
	}
}

/*
Sort - returns the sorted slice
*/
func (receiver *ArraySorting) Sort(array []string) []string {
	var sortedArray [][]string

	regExpr := regexp.MustCompile(`\S+\s*`)

	var funcSplit func(key int, value string)
	uniqueValuesMap := make(map[string]struct{})

	if receiver.orderFlagDTO.UniqueFlag {
		funcSplit = func(key int, value string) {
			tempArray := receiver.splitAndTrimRow(regExpr, value)
			tempValue := tempArray[receiver.workingColumNumber]

			if _, isKeyExists := uniqueValuesMap[tempValue]; !isKeyExists {
				uniqueValuesMap[tempValue] = struct{}{}
				sortedArray = append(sortedArray, tempArray)
			}
		}
	} else {
		funcSplit = func(key int, value string) {
			sortedArray = append(sortedArray, receiver.splitAndTrimRow(regExpr, value))
		}
	}

	for key, value := range array {
		funcSplit(key, value)
	}

	if receiver.orderFlagDTO.SortCheckFlag {
		isSorted := slices.IsSortedFunc(sortedArray, receiver.runSortFunc)
		return []string{strconv.FormatBool(isSorted)}
	}

	slices.SortFunc(sortedArray, receiver.runSortFunc)

	return receiver.joinArray(sortedArray)
}

func (receiver *ArraySorting) splitAndTrimRow(regexp *regexp.Regexp, row string) []string {
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

func (receiver *ArraySorting) runSortFunc(a, b []string) int {
	firstValue := strings.ToLower(a[receiver.workingColumNumber])
	secondValue := strings.ToLower(b[receiver.workingColumNumber])

	sortMethod := NewSortMethod(receiver.orderFlagDTO)

	if result := sortMethod.Execute(firstValue, secondValue); result != 0 {
		if receiver.orderFlagDTO.DescendingOrderFlag {
			return -1 * result
		}
		return result
	}

	if receiver.orderFlagDTO.DescendingOrderFlag {
		return strings.Compare(a[receiver.workingColumNumber], b[receiver.workingColumNumber])
	}

	return strings.Compare(b[receiver.workingColumNumber], a[receiver.workingColumNumber])
}

func (receiver *ArraySorting) joinArray(sortedArray [][]string) []string {
	joinedArray := make([]string, len(sortedArray))

	for index, value := range sortedArray {
		joinedArray[index] = strings.Join(value, " ")
	}

	return joinedArray
}
