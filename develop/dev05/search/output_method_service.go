package search

type OutputMethodService struct {
	inputRows         []string
	foundRowMap       map[int]string
	outputSliceLength int
}

func NewOutputMethodService(inputRows []string, foundRowMap map[int]string) *OutputMethodService {
	return &OutputMethodService{
		inputRows:         inputRows,
		foundRowMap:       foundRowMap,
		outputSliceLength: len(inputRows),
	}
}

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
