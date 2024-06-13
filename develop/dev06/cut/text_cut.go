package cut

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TextCut struct {
	dto *TextCutFlagDTO
}

func NewTextCut(dto *TextCutFlagDTO) *TextCut {
	return &TextCut{dto}
}

func (receiver *TextCut) Cut(inputRows []string) []string {
	fields := receiver.convertValuesOfFieldsFlag()

	var outputRows []string

	for _, inputRow := range inputRows {
		splitRows := strings.Split(inputRow, receiver.dto.DelimiterFlag)

		var outputRow string

		if len(splitRows) > 1 {
			for index, fieldNumber := range fields {
				if index == 0 {
					outputRow = splitRows[fieldNumber]
					continue
				}
				outputRow += receiver.dto.DelimiterFlag + splitRows[fieldNumber]
			}
			outputRows = append(outputRows, outputRow)
			continue
		}

		if !receiver.dto.SeparatedFlag {
			outputRows = append(outputRows, inputRow)
		}
	}

	return outputRows
}

func (receiver *TextCut) convertValuesOfFieldsFlag() []int {
	stringArray := strings.Split(receiver.dto.FieldsFlag, ",")

	var integerArray []int

	for _, value := range stringArray {
		number, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("parameter -f error")
			os.Exit(1)
		}

		integerArray = append(integerArray, number-1)
	}

	return integerArray
}
