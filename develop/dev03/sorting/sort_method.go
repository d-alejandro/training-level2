package sorting

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

/*
SortMethod structure - sorting by characteristics
*/
type SortMethod struct {
	orderFlagDTO *FlagDTO
}

/*
NewSortMethod is the SortMethod constructor
*/
func NewSortMethod(orderFlagDTO *FlagDTO) *SortMethod {
	return &SortMethod{orderFlagDTO}
}

/*
Execute - returns the sort result
*/
func (receiver *SortMethod) Execute(firstValue, secondValue string) int {
	if firstValue == secondValue {
		return 0
	}

	switch {
	case receiver.orderFlagDTO.SortByNumberFlag:
		if response, err := receiver.sortByNumber(firstValue, secondValue); err == nil {
			return response
		}
	case receiver.orderFlagDTO.MonthNameFlag:
		if response, err := receiver.sortByMonth(firstValue, secondValue); err == nil {
			return response
		}
	case receiver.orderFlagDTO.HumanNumericFlag:
		if response, err := receiver.sortByHumanReadableSize(firstValue, secondValue); err == nil {
			return response
		}
	}

	return strings.Compare(firstValue, secondValue)
}

func (receiver *SortMethod) sortByNumber(firstValue, secondValue string) (int, error) {
	var (
		firstValueInt  int
		secondValueInt int
		err            error
	)

	if firstValueInt, err = strconv.Atoi(firstValue); err != nil {
		return 0, err
	}

	if secondValueInt, err = strconv.Atoi(secondValue); err != nil {
		return 0, err
	}

	if firstValueInt < secondValueInt {
		return -1, nil
	}

	return 1, nil
}

func (receiver *SortMethod) sortByMonth(firstValue, secondValue string) (int, error) {
	var (
		firstValueMonth  time.Time
		secondValueMonth time.Time
		err              error
	)

	if firstValueMonth, err = time.Parse("Jan", firstValue); err != nil {
		return 0, err
	}

	if secondValueMonth, err = time.Parse("Jan", secondValue); err != nil {
		return 0, err
	}

	return firstValueMonth.Compare(secondValueMonth), nil
}

func (receiver *SortMethod) sortByHumanReadableSize(firstValue, secondValue string) (int, error) {
	var (
		firstValueSuffix   string
		secondValueSuffix  string
		firstValueString   string
		secondValueString  string
		firstSuffixNumber  int
		secondSuffixNumber int
		isKeyExists        bool
	)

	suffixes := map[string]int{
		"b": 1,
		"k": 2,
		"m": 3,
		"g": 4,
		"t": 5,
		"p": 6,
		"e": 7,
	}

	firstValueLength := len(firstValue)
	firstValueSuffix = firstValue[firstValueLength-1:]

	if firstSuffixNumber, isKeyExists = suffixes[firstValueSuffix]; !isKeyExists {
		return 0, errors.New("first suffix number not found")
	}

	firstValueString = firstValue[:firstValueLength-1]

	secondValueLength := len(secondValue)
	secondValueSuffix = secondValue[secondValueLength-1:]

	if secondSuffixNumber, isKeyExists = suffixes[secondValueSuffix]; !isKeyExists {
		return 0, errors.New("second suffix number not found")
	}

	secondValueString = secondValue[:secondValueLength-1]

	if firstSuffixNumber == secondSuffixNumber {
		return receiver.sortByNumber(firstValueString, secondValueString)
	}

	if firstSuffixNumber < secondSuffixNumber {
		return -1, nil
	}

	return 1, nil
}
