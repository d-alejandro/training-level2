package search

import "slices"

/*
OutputMethodService structure
*/
type OutputMethodService struct {
	inputRows         []string
	foundRowMap       map[int]string
	outputSliceLength int
}

/*
NewOutputMethodService constructor
*/
func NewOutputMethodService(inputRows []string, foundRowMap map[int]string) *OutputMethodService {
	return &OutputMethodService{
		inputRows:         inputRows,
		foundRowMap:       foundRowMap,
		outputSliceLength: len(inputRows),
	}
}

/*
ExecuteForRowsAfterFlag method
*/
func (receiver *OutputMethodService) ExecuteForRowsAfterFlag(rowCountAfter int) []string {
	outputSlice := make([]string, receiver.outputSliceLength)

	for key := range outputSlice {
		if value, ok := receiver.foundRowMap[key]; ok {
			outputSlice[key] = value

			lastIndex := key + 1 + rowCountAfter

			if lastIndex > receiver.outputSliceLength {
				lastIndex = receiver.outputSliceLength
			}

			for index := key + 1; index < lastIndex; index++ {
				outputSlice[index] = receiver.inputRows[index]
			}
		}
	}

	return outputSlice
}

/*
ExecuteForRowsBeforeFlag method
*/
func (receiver *OutputMethodService) ExecuteForRowsBeforeFlag(rowCountBefore int) []string {
	outputSlice := make([]string, receiver.outputSliceLength)

	for key := receiver.outputSliceLength - 1; key >= 0; key-- {
		if value, ok := receiver.foundRowMap[key]; ok {
			outputSlice[key] = value

			lastIndex := key - rowCountBefore

			if lastIndex < 0 {
				lastIndex = 0
			}

			for index := key - 1; index >= lastIndex; index-- {
				outputSlice[index] = receiver.inputRows[index]
			}
		}
	}

	return outputSlice
}

/*
ExecuteForRowsContextFlag method
*/
func (receiver *OutputMethodService) ExecuteForRowsContextFlag(rowContext int) []string {
	outputSliceFirst := receiver.ExecuteForRowsAfterFlag(rowContext)
	outputSliceSecond := receiver.ExecuteForRowsBeforeFlag(rowContext)

	for key, value := range outputSliceFirst {
		if value == "" {
			outputSliceFirst[key] = outputSliceSecond[key]
		}
	}

	return outputSliceFirst
}

/*
ExecuteWithoutFlags method
*/
func (receiver *OutputMethodService) ExecuteWithoutFlags() []string {
	outputSlice := make([]string, receiver.outputSliceLength)

	for key, value := range receiver.foundRowMap {
		outputSlice[key] = value
	}

	return outputSlice
}

/*
ExecuteForInvertFlag method
*/
func (receiver *OutputMethodService) ExecuteForInvertFlag() []string {
	outputSlice := slices.Clone(receiver.inputRows)

	for key := range receiver.foundRowMap {
		outputSlice[key] = ""
	}

	return outputSlice
}
