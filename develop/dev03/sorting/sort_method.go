package sorting

import (
	"strconv"
	"strings"
	"time"
)

type SortMethod struct {
	orderFlagDTO *FlagDTO
}

func NewSortMethod(orderFlagDTO *FlagDTO) *SortMethod {
	return &SortMethod{orderFlagDTO}
}

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
