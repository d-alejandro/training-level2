package sorting

import (
	"strconv"
	"strings"
)

type SortMethod struct {
	orderFlagDTO *FlagDTO
}

func NewSortMethod(orderFlagDTO *FlagDTO) *SortMethod {
	return &SortMethod{orderFlagDTO}
}

func (receiver *SortMethod) Execute(firstValue, secondValue string) int {
	switch {
	case receiver.orderFlagDTO.SortByNumberFlag:
		if response, err := receiver.sortByNumber(firstValue, secondValue); err == nil {
			return response
		}
	}

	return strings.Compare(firstValue, secondValue)
}

func (receiver *SortMethod) sortByNumber(firstValue, secondValue string) (int, error) {
	if firstValue == secondValue {
		return 0, nil
	}

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
